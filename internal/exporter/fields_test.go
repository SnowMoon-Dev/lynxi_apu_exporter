package exporter

import (
  _ "embed"
  "os/exec"
  "testing"

  "github.com/stretchr/testify/assert"
)

var (
  //go:embed _fields-test.txt
  fieldsTest      string
  expectedQFields = []qField{
    __TIMESTAMP_KEY__,
    __BOARD_INDEX_KEY__,
    __PRODUCT_NAME_KEY__,
    __PRODUCT_NUMBER_KEY__,
    __DRIVER_VERSION_KEY__,
    __FIRMWARE_VERSION_KEY__,
    __SERIAL_NUMBER_KEY__,
    __CHIP_COUNT_KEY__,
    __CHIP_ID_CHIP0_KEY__,
    __CHIP_ID_CHIP1_KEY__,
    __CHIP_ID_CHIP2_KEY__,
    __UUID_CHIP0_KEY__,
    __UUID_CHIP1_KEY__,
    __UUID_CHIP2_KEY__,
    __CHIP_INDEX_CHIP0_KEY__,
    __CHIP_INDEX_CHIP1_KEY__,
    __CHIP_INDEX_CHIP2_KEY__,
    __UTILIZATION_APU_TOTAL_KEY__,
    __UTILIZATION_APU_CHIP0_KEY__,
    __UTILIZATION_APU_CHIP1_KEY__,
    __UTILIZATION_APU_CHIP2_KEY__,
    __UTILIZATION_CPU_TOTAL_KEY__,
    __UTILIZATION_CPU_CHIP0_KEY__,
    __UTILIZATION_CPU_CHIP1_KEY__,
    __UTILIZATION_CPU_CHIP2_KEY__,
    __UTILIZATION_VIC_TOTAL_KEY__,
    __UTILIZATION_VIC_CHIP0_KEY__,
    __UTILIZATION_VIC_CHIP1_KEY__,
    __UTILIZATION_VIC_CHIP2_KEY__,
    __UTILIZATION_MEMORY_TOTAL_KEY__,
    __UTILIZATION_MEMORY_CHIP0_KEY__,
    __UTILIZATION_MEMORY_CHIP1_KEY__,
    __UTILIZATION_MEMORY_CHIP2_KEY__,
    __UTILIZATION_IPE_FPS_TOTAL_KEY__,
    __UTILIZATION_IPE_FPS_CHIP0_KEY__,
    __UTILIZATION_IPE_FPS_CHIP1_KEY__,
    __UTILIZATION_IPE_FPS_CHIP2_KEY__,
    __PCI_SUB_VENDOR_ID_CHIP0_KEY__,
    __PCI_SUB_VENDOR_ID_CHIP1_KEY__,
    __PCI_SUB_VENDOR_ID_CHIP2_KEY__,
    __PCI_VENDOR_ID_CHIP0_KEY__,
    __PCI_VENDOR_ID_CHIP1_KEY__,
    __PCI_VENDOR_ID_CHIP2_KEY__,
    __PCI_BUS_CHIP0_KEY__,
    __PCI_BUS_CHIP1_KEY__,
    __PCI_BUS_CHIP2_KEY__,
    __PCI_DEVICE_ID_CHIP0_KEY__,
    __PCI_DEVICE_ID_CHIP1_KEY__,
    __PCI_DEVICE_ID_CHIP2_KEY__,
    __PCI_SUB_DEVICE_ID_CHIP0_KEY__,
    __PCI_SUB_DEVICE_ID_CHIP1_KEY__,
    __PCI_SUB_DEVICE_ID_CHIP2_KEY__,
    __PCI_DEVICE_CHIP0_KEY__,
    __PCI_DEVICE_CHIP1_KEY__,
    __PCI_DEVICE_CHIP2_KEY__,
    __PCI_FUNCTION_CHIP0_KEY__,
    __PCI_FUNCTION_CHIP1_KEY__,
    __PCI_FUNCTION_CHIP2_KEY__,
    __PCI_NUMA_NODE_ID_CHIP0_KEY__,
    __PCI_NUMA_NODE_ID_CHIP1_KEY__,
    __PCI_NUMA_NODE_ID_CHIP2_KEY__,
    __PCI_NUMA_CPU_CHIP0_KEY__,
    __PCI_NUMA_CPU_CHIP1_KEY__,
    __PCI_NUMA_CPU_CHIP2_KEY__,
    __PCIE_LINK_SPEED_MAX_CHIP0_KEY__,
    __PCIE_LINK_SPEED_MAX_CHIP1_KEY__,
    __PCIE_LINK_SPEED_MAX_CHIP2_KEY__,
    __PCIE_LINK_SPEED_CURRENT_CHIP0_KEY__,
    __PCIE_LINK_SPEED_CURRENT_CHIP1_KEY__,
    __PCIE_LINK_SPEED_CURRENT_CHIP2_KEY__,
    __PCIE_LINK_GEN_MAX_CHIP0_KEY__,
    __PCIE_LINK_GEN_MAX_CHIP1_KEY__,
    __PCIE_LINK_GEN_MAX_CHIP2_KEY__,
    __PCIE_LINK_GEN_CURRENT_CHIP0_KEY__,
    __PCIE_LINK_GEN_CURRENT_CHIP1_KEY__,
    __PCIE_LINK_GEN_CURRENT_CHIP2_KEY__,
    __FAN_SPEED_KEY__,
    __TEMPERATURE_CURRENT_CHIP0_KEY__,
    __TEMPERATURE_CURRENT_CHIP1_KEY__,
    __TEMPERATURE_CURRENT_CHIP2_KEY__,
    __VOLTAGE_CURRENT_CHIP0_KEY__,
    __VOLTAGE_CURRENT_CHIP1_KEY__,
    __VOLTAGE_CURRENT_CHIP2_KEY__,
    __VOLTAGE_BOARD_INPUT_KEY__,
    __CLOCKS_CURRENT_APU_CHIP0_KEY__,
    __CLOCKS_CURRENT_APU_CHIP1_KEY__,
    __CLOCKS_CURRENT_APU_CHIP2_KEY__,
    __CLOCKS_CURRENT_CPU_CHIP0_KEY__,
    __CLOCKS_CURRENT_CPU_CHIP1_KEY__,
    __CLOCKS_CURRENT_CPU_CHIP2_KEY__,
    __CLOCKS_CURRENT_MEMORY_CHIP0_KEY__,
    __CLOCKS_CURRENT_MEMORY_CHIP1_KEY__,
    __CLOCKS_CURRENT_MEMORY_CHIP2_KEY__,
    __CLOCKS_MAX_APU_CHIP0_KEY__,
    __CLOCKS_MAX_APU_CHIP1_KEY__,
    __CLOCKS_MAX_APU_CHIP2_KEY__,
    __CLOCKS_MAX_CPU_CHIP0_KEY__,
    __CLOCKS_MAX_CPU_CHIP1_KEY__,
    __CLOCKS_MAX_CPU_CHIP2_KEY__,
    __CLOCKS_MAX_MEMORY_CHIP0_KEY__,
    __CLOCKS_MAX_MEMORY_CHIP1_KEY__,
    __CLOCKS_MAX_MEMORY_CHIP2_KEY__,
    __POWER_DRAW__,
    __POWER_LIMIT__,
    __ECC_MODE_CURRENT_CHIP0_KEY__,
    __ECC_MODE_CURRENT_CHIP1_KEY__,
    __ECC_MODE_CURRENT_CHIP2_KEY__,
    __ECC_ERRORS_CORRECTED_TOTAL_KEY__,
    __ECC_ERRORS_CORRECTED_TOTAL_CHIP0_KEY__,
    __ECC_ERRORS_CORRECTED_TOTAL_CHIP1_KEY__,
    __ECC_ERRORS_CORRECTED_TOTAL_CHIP2_KEY__,
    __ECC_ERRORS_UNCORRECTED_TOTAL_KEY__,
    __ECC_ERRORS_UNCORRECTED_TOTAL_CHIP0_KEY__,
    __ECC_ERRORS_UNCORRECTED_TOTAL_CHIP1_KEY__,
    __ECC_ERRORS_UNCORRECTED_TOTAL_CHIP2_KEY__,
  }
)

func TestExtractQFields(t *testing.T) {
  fields := extractQFields(fieldsTest)
  assert.Equal(t, expectedQFields, fields)
}

func TestParseAutoQFields(t *testing.T) {
  runCmdOriginal := runCmd
  defer func() { runCmd = runCmdOriginal }()

  var capturedCmd *exec.Cmd
  runCmd = func(cmd *exec.Cmd) error {
    capturedCmd = cmd
    _, _ = cmd.Stdout.Write([]byte(fieldsTest))
    return nil
  }

  fields, err := ParseAutoQFields("lynxi-smi-pro")

  assert.Len(t, capturedCmd.Args, 2)
  assert.Equal(t, capturedCmd.Args[0], "lynxi-smi-pro")
  assert.Equal(t, capturedCmd.Args[1], "--help-query-apu")

  assert.NoError(t, err)
  assert.Equal(t, expectedQFields, fields)
}
