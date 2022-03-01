package exporter

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// qField stands for query field - the field name before the query
type qField string

// rField stands for returned field - the field name as returned by the lynxi-smi
type rField string

const (
	DefaultPrefix           = "lynxi_smi_pro"
	DefaultLynSmiProCommand = "lynxi-smi-pro"
)

var (
	numericRegex = regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+")

	requiredFields = []requiredField{
		{qField: nameQField, label: "name"},
		{qField: boardIndexQField, label: "board_index"},
		{qField: productNumberQField, label: "product_number"},
		{qField: serialNumberQField, label: "serial_number"},
		{qField: chipCountQField, label: "chip_count"},
		{qField: driverVersionQField, label: "driver_version"},
		{qField: firmwareVersionQField, label: "firmware_version"},
	}

	runCmd = func(cmd *exec.Cmd) error { return cmd.Run() }
)

// Exporter collects stats and exports them using
// the prometheus metrics package.
type gpuExporter struct {
	mutex                 sync.RWMutex
	prefix                string
	qFields               []qField
	qFieldToMetricInfoMap map[qField]MetricInfo
	lynSmiProCommand      string
	failedScrapesTotal    prometheus.Counter
	exitCode              prometheus.Gauge
	apuInfoDesc           *prometheus.Desc
	logger                log.Logger
}

type MetricInfo struct {
	desc            *prometheus.Desc
	mType           prometheus.ValueType
	valueMultiplier float64
}

type requiredField struct {
	qField qField
	label  string
}

func New(prefix string, lynSmiProCommand string, qFieldsRaw string, logger log.Logger) (prometheus.Collector, error) {

	qFieldsOrdered, qFieldToRFieldMap, err := buildQFieldToRFieldMap(logger, qFieldsRaw, lynSmiProCommand)
	if err != nil {
		return nil, err
	}

	qFieldToMetricInfoMap := buildQFieldToMetricInfoMap(prefix, qFieldToRFieldMap)

	infoLabels := getLabels(requiredFields)
	e := gpuExporter{
		prefix:                prefix,
		lynSmiProCommand:      lynSmiProCommand,
		qFields:               qFieldsOrdered,
		qFieldToMetricInfoMap: qFieldToMetricInfoMap,
		logger:                logger,
		failedScrapesTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: prefix,
			Name:      "failed_scrapes_total",
			Help:      "Number of failed scrapes",
		}),
		exitCode: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: prefix,
			Name:      "command_exit_code",
			Help:      "Exit code of the last scrape command",
		}),
		apuInfoDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prefix, "", "apu_info"),
			fmt.Sprintf("A metric with a constant '1' value labeled by apu %s.",
				strings.Join(infoLabels, ", ")),
			infoLabels,
			nil),
	}

	return &e, nil
}

// Describe describes all the metrics ever exported by the exporter. It
// implements prometheus.Collector.
func (e *gpuExporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range e.qFieldToMetricInfoMap {
		ch <- m.desc
	}
	ch <- e.failedScrapesTotal.Desc()
	ch <- e.apuInfoDesc
}

// Collect fetches the stats and delivers them as Prometheus metrics. It implements prometheus.Collector.
func (e *gpuExporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	exitCode, t, err := scrape(e.qFields, e.lynSmiProCommand)
	e.exitCode.Set(float64(exitCode))
	ch <- e.exitCode
	if err != nil {
		_ = level.Error(e.logger).Log("error", err)
		ch <- e.failedScrapesTotal
		e.failedScrapesTotal.Inc()
		return
	}

	for _, r := range t.rows {
		name := r.qFieldToCells[nameQField].rawValue
		boardIndex := r.qFieldToCells[boardIndexQField].rawValue
		productNumber := r.qFieldToCells[productNumberQField].rawValue
		serialNumber := r.qFieldToCells[serialNumberQField].rawValue
		chipCount := r.qFieldToCells[chipCountQField].rawValue
		driverVersion := r.qFieldToCells[driverVersionQField].rawValue
		firmwareVersion := r.qFieldToCells[firmwareVersionQField].rawValue
		infoMetric := prometheus.MustNewConstMetric(e.apuInfoDesc, prometheus.GaugeValue,
			1, name, boardIndex, productNumber,
			serialNumber, chipCount, driverVersion, firmwareVersion)
		ch <- infoMetric
		for _, c := range r.cells {
			mi := e.qFieldToMetricInfoMap[c.qField]
			num, err := transformRawValue(c.rawValue, mi.valueMultiplier)
			if err != nil {
				_ = level.Debug(e.logger).Log("error", err, "query_field_name",
					c.qField, "raw_value", c.rawValue)
				continue
			}
			ch <- prometheus.MustNewConstMetric(mi.desc, mi.mType, num, name, boardIndex, serialNumber, firmwareVersion)
		}
	}
}

func transformRawValue(rawValue string, valueMultiplier float64) (float64, error) {
	trimmed := strings.TrimSpace(rawValue)
	if strings.HasPrefix(trimmed, "0x") {
		return hexToDecimal(trimmed)
	}

	val := strings.ToLower(trimmed)
	switch val {
	case "enabled", "yes", "active":
		return 1, nil
	case "disabled", "no", "not active":
		return 0, nil
	case "default":
		return 0, nil
	case "exclusive_thread":
		return 1, nil
	case "prohibited":
		return 2, nil
	case "exclusive_process":
		return 3, nil
	default:
		allNums := numericRegex.FindAllString(val, 2)
		if len(allNums) != 1 {
			return -1, fmt.Errorf("couldn't parse number from: %s", val)
		}

		parsed, err := strconv.ParseFloat(allNums[0], 64)
		if err != nil {
			return -1, err
		}
		return parsed * valueMultiplier, err
	}
}

func getLabels(reqFields []requiredField) []string {
	r := make([]string, len(reqFields))
	for i, reqField := range reqFields {
		r[i] = reqField.label
	}
	return r
}

func buildQFieldToRFieldMap(logger log.Logger, qFieldsRaw string,
	LynSmiProCommand string) ([]qField, map[qField]rField, error) {
	qFieldsSeparated := strings.Split(qFieldsRaw, __COMMA_SEP__)

	qFields := toQFieldSlice(qFieldsSeparated)
	for _, reqField := range requiredFields {
		qFields = append(qFields, reqField.qField)
	}

	qFields = removeDuplicateQFields(qFields)

	if len(qFieldsSeparated) == 1 && qFieldsSeparated[0] == qFieldsAuto {
		parsed, err := ParseAutoQFields(LynSmiProCommand)
		if err != nil {
			_ = level.Warn(logger).Log("msg",
				"Failed to auto-determine query field names, "+
					"falling back to the built-in list")
			return getKeys(fallbackQFieldToRFieldMap), fallbackQFieldToRFieldMap, nil
		}

		qFields = parsed
	}

	_, t, err := scrape(qFields, LynSmiProCommand)
	var rFields []rField
	if err != nil {
		_ = level.Warn(logger).Log("msg",
			"Failed to run an initial scrape, using the built-in list for field mapping")
		rFields, err = getFallbackValues(qFields)
		if err != nil {
			return nil, nil, err
		}
	} else {
		rFields = t.rFields
	}

	r := make(map[qField]rField, len(qFields))
	for i, q := range qFields {
		r[q] = rFields[i]
	}

	return qFields, r, nil
}

func getFallbackValues(qFields []qField) ([]rField, error) {
	r := make([]rField, len(qFields))
	i := 0
	for _, q := range qFields {
		val, contains := fallbackQFieldToRFieldMap[q]
		if !contains {
			return nil, fmt.Errorf("unexpected query field: %s", q)
		}
		r[i] = val
		i++
	}
	return r, nil
}

func getKeys(m map[qField]rField) []qField {
	r := make([]qField, len(m))
	i := 0
	for key := range m {
		r[i] = key
		i++
	}
	return r
}

func scrape(qFields []qField, lynSmiProCommand string) (int, *table, error) {
	qFieldsJoined := strings.Join(QFieldSliceToStringSlice(qFields), __COMMA_SEP__)
	cmdAndArgs := strings.Fields(lynSmiProCommand)

	cmdAndArgs = append(cmdAndArgs, fmt.Sprintf("--query-apu=%s", qFieldsJoined))

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := runCmd(cmd)
	if err != nil {
		exitCode := -1
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		}
		return exitCode, nil, fmt.Errorf("command failed. stderr: %s err: %w", stderr.String(), err)
	}

	t, err := parseCSVIntoTable(strings.TrimSpace(stdout.String()), qFields)
	if err != nil {
		return -1, nil, err
	}

	return 0, &t, nil
}
func toQFieldSlice(ss []string) []qField {
	r := make([]qField, len(ss))
	for i, s := range ss {
		r[i] = qField(s)
	}
	return r
}

func toRFieldSlice(ss []string) []rField {
	r := make([]rField, len(ss))
	for i, s := range ss {
		r[i] = rField(s)
	}
	return r
}

func QFieldSliceToStringSlice(qs []qField) []string {
	r := make([]string, len(qs))
	for i, q := range qs {
		r[i] = string(q)
	}
	return r
}

func extractQFields(text string) []qField {
	found := fieldRegex.FindAllStringSubmatch(text, -1)

	var fields []qField
	for _, ss := range found {
		fields = append(fields, qField(ss[1]))
	}
	return fields
}

func removeDuplicateQFields(qFields []qField) []qField {
	m := make(map[qField]struct{})
	var r []qField
	for _, f := range qFields {
		_, exists := m[f]
		if !exists {
			r = append(r, f)
			m[f] = struct{}{}
		}
	}
	return r
}
func ParseAutoQFields(lynSmiProCommand string) ([]qField, error) {
	cmdAndArgs := strings.Fields(lynSmiProCommand)
	cmdAndArgs = append(cmdAndArgs, "--help-query-apu")
	cmd := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := runCmd(cmd)
	if err != nil {
		return nil, err
	}

	out := stdout.String()
	fields := extractQFields(out)
	if fields == nil {
		return nil, errors.New("could not extract any query fields")
	}
	return fields, nil
}

func buildQFieldToMetricInfoMap(prefix string, qFieldToRFieldMap map[qField]rField) map[qField]MetricInfo {
	result := make(map[qField]MetricInfo)
	for qField, rField := range qFieldToRFieldMap {
		result[qField] = buildMetricInfo(prefix, rField)
	}
	return result
}

func buildMetricInfo(prefix string, rField rField) MetricInfo {
	fqName, multiplier := buildFQNameAndMultiplier(prefix, rField)
	desc := prometheus.NewDesc(fqName, string(rField), []string{"name", "board_index", "serial_number", "firmware_version"}, nil)
	return MetricInfo{
		desc:            desc,
		mType:           prometheus.GaugeValue,
		valueMultiplier: multiplier,
	}
}

func buildFQNameAndMultiplier(prefix string, rField rField) (string, float64) {
	rFieldStr := string(rField)
	suffixTransformed := rFieldStr
	multiplier := 1.0
	split := strings.Split(rFieldStr, " ")[0]
	if strings.HasSuffix(rFieldStr, " [W]") {
		suffixTransformed = split + "_watts"
	} else if strings.HasSuffix(rFieldStr, " [MHz]") {
		suffixTransformed = split + "_clock_hz"
		multiplier = 1000000
	} else if strings.HasSuffix(rFieldStr, " [MiB]") {
		suffixTransformed = split + "_bytes"
		multiplier = 1048576
	} else if strings.HasSuffix(rFieldStr, " [%]") {
		suffixTransformed = split + "_ratio"
		multiplier = 0.01
	}

	metricName := toSnakeCase(strings.ReplaceAll(suffixTransformed, ".", "_"))
	fqName := prometheus.BuildFQName(prefix, "", metricName)
	return fqName, multiplier
}
