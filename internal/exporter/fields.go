package exporter

import (
	"regexp"
)

const (
	nameQField            qField = "name"
	boardIndexQField      qField = "board_index"
	productNumberQField   qField = "product_number"
	serialNumberQField    qField = "serial_number"
	chipCountQField       qField = "chip_count"
	driverVersionQField   qField = "driver_version"
	firmwareVersionQField qField = "firmware_version"
	qFieldsAuto                  = "AUTO"
	DefaultQField                = qFieldsAuto
)

const (
	__COLON_SEP__     = ":"
	__LINE_FEED_STR__ = "\n"
	__LINE_FEED_SEP__ = '\n'
	__PER_SEP__       = "%"
	__V_SEP__         = "V"
	__C_SEP__         = "C"
	__SPCAE_SEP__     = " "
	__W_SEP__         = "W"
	__DOT_SEP__       = "."
	__COMMA_SEP__     = ","
)

var (
	fieldRegex     = regexp.MustCompile(`(?m)\n\s*\n^"([^"]+)"`)
	fallbackQField = []qField{
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
		__PCI_VENDOR_ID_CHIP0_KEY__,
		__PCI_VENDOR_ID_CHIP1_KEY__,
		__PCI_VENDOR_ID_CHIP2_KEY__,
		__PCI_DEVICE_ID_CHIP0_KEY__,
		__PCI_DEVICE_ID_CHIP1_KEY__,
		__PCI_DEVICE_ID_CHIP2_KEY__,
		__PCI_SUB_VENDOR_ID_CHIP0_KEY__,
		__PCI_SUB_VENDOR_ID_CHIP1_KEY__,
		__PCI_SUB_VENDOR_ID_CHIP2_KEY__,
		__PCI_SUB_DEVICE_ID_CHIP0_KEY__,
		__PCI_SUB_DEVICE_ID_CHIP1_KEY__,
		__PCI_SUB_DEVICE_ID_CHIP2_KEY__,
		__PCI_BUS_CHIP0_KEY__,
		__PCI_BUS_CHIP1_KEY__,
		__PCI_BUS_CHIP2_KEY__,
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
	fallbackQFieldToRFieldMap = map[qField]rField{
		__TIMESTAMP_KEY__:                          "timestamp",
		__BOARD_INDEX_KEY__:                        "board_index",
		__PRODUCT_NAME_KEY__:                       "name",
		__PRODUCT_NUMBER_KEY__:                     "product_number",
		__DRIVER_VERSION_KEY__:                     "driver_version",
		__FIRMWARE_VERSION_KEY__:                   "firmware_version",
		__SERIAL_NUMBER_KEY__:                      "serial_number",
		__CHIP_COUNT_KEY__:                         "chip_count",
		__CHIP_ID_CHIP0_KEY__:                      "chip_id.chip0",
		__CHIP_ID_CHIP1_KEY__:                      "chip_id.chip1",
		__CHIP_ID_CHIP2_KEY__:                      "chip_id.chip2",
		__UUID_CHIP0_KEY__:                         "uuid.chip0",
		__UUID_CHIP1_KEY__:                         "uuid.chip1",
		__UUID_CHIP2_KEY__:                         "uuid.chip2",
		__CHIP_INDEX_CHIP0_KEY__:                   "chip_index.chip0",
		__CHIP_INDEX_CHIP1_KEY__:                   "chip_index.chip1",
		__CHIP_INDEX_CHIP2_KEY__:                   "chip_index.chip2",
		__UTILIZATION_APU_TOTAL_KEY__:              "utilization.apu.total",
		__UTILIZATION_APU_CHIP0_KEY__:              "utilization.apu.chip0",
		__UTILIZATION_APU_CHIP1_KEY__:              "utilization.apu.chip1",
		__UTILIZATION_APU_CHIP2_KEY__:              "utilization.apu.chip2",
		__UTILIZATION_CPU_TOTAL_KEY__:              "utilization.cpu.total",
		__UTILIZATION_CPU_CHIP0_KEY__:              "utilization.cpu.chip0",
		__UTILIZATION_CPU_CHIP1_KEY__:              "utilization.cpu.chip1",
		__UTILIZATION_CPU_CHIP2_KEY__:              "utilization.cpu.chip2",
		__UTILIZATION_VIC_TOTAL_KEY__:              "utilization.vic.total",
		__UTILIZATION_VIC_CHIP0_KEY__:              "utilization.vic.chip0",
		__UTILIZATION_VIC_CHIP1_KEY__:              "utilization.vic.chip1",
		__UTILIZATION_VIC_CHIP2_KEY__:              "utilization.vic.chip2",
		__UTILIZATION_MEMORY_TOTAL_KEY__:           "utilization.memory.total",
		__UTILIZATION_MEMORY_CHIP0_KEY__:           "utilization.memory.chip0",
		__UTILIZATION_MEMORY_CHIP1_KEY__:           "utilization.memory.chip1",
		__UTILIZATION_MEMORY_CHIP2_KEY__:           "utilization.memory.chip2",
		__UTILIZATION_IPE_FPS_TOTAL_KEY__:          "utilization.ipeFps.total",
		__UTILIZATION_IPE_FPS_CHIP0_KEY__:          "utilization.ipeFps.chip0",
		__UTILIZATION_IPE_FPS_CHIP1_KEY__:          "utilization.ipeFps.chip1",
		__UTILIZATION_IPE_FPS_CHIP2_KEY__:          "utilization.ipeFps.chip2",
		__PCI_SUB_VENDOR_ID_CHIP0_KEY__:            "pci.sub_vendor_id.chip0",
		__PCI_SUB_VENDOR_ID_CHIP1_KEY__:            "pci.sub_vendor_id.chip1",
		__PCI_SUB_VENDOR_ID_CHIP2_KEY__:            "pci.sub_vendor_id.chip2",
		__PCI_VENDOR_ID_CHIP0_KEY__:                "pci.vendor_id.chip0",
		__PCI_VENDOR_ID_CHIP1_KEY__:                "pci.vendor_id.chip1",
		__PCI_VENDOR_ID_CHIP2_KEY__:                "pci.vendor_id.chip2",
		__PCI_BUS_CHIP0_KEY__:                      "pci.bus.chip0",
		__PCI_BUS_CHIP1_KEY__:                      "pci.bus.chip1",
		__PCI_BUS_CHIP2_KEY__:                      "pci.bus.chip2",
		__PCI_DEVICE_ID_CHIP0_KEY__:                "pci.device_id.chip0",
		__PCI_DEVICE_ID_CHIP1_KEY__:                "pci.device_id.chip1",
		__PCI_DEVICE_ID_CHIP2_KEY__:                "pci.device_id.chip2",
		__PCI_SUB_DEVICE_ID_CHIP0_KEY__:            "pci.sub_device_id.chip0",
		__PCI_SUB_DEVICE_ID_CHIP1_KEY__:            "pci.sub_device_id.chip1",
		__PCI_SUB_DEVICE_ID_CHIP2_KEY__:            "pci.sub_device_id.chip2",
		__PCI_DEVICE_CHIP0_KEY__:                   "pci.device.chip0",
		__PCI_DEVICE_CHIP1_KEY__:                   "pci.device.chip1",
		__PCI_DEVICE_CHIP2_KEY__:                   "pci.device.chip2",
		__PCI_FUNCTION_CHIP0_KEY__:                 "pci.function.chip0",
		__PCI_FUNCTION_CHIP1_KEY__:                 "pci.function.chip1",
		__PCI_FUNCTION_CHIP2_KEY__:                 "pci.function.chip2",
		__PCI_NUMA_NODE_ID_CHIP0_KEY__:             "pci.numa.node_id.chip0",
		__PCI_NUMA_NODE_ID_CHIP1_KEY__:             "pci.numa.node_id.chip1",
		__PCI_NUMA_NODE_ID_CHIP2_KEY__:             "pci.numa.node_id.chip2",
		__PCI_NUMA_CPU_CHIP0_KEY__:                 "pci.numa.cpu.chip0",
		__PCI_NUMA_CPU_CHIP1_KEY__:                 "pci.numa.cpu.chip1",
		__PCI_NUMA_CPU_CHIP2_KEY__:                 "pci.numa.cpu.chip2",
		__PCIE_LINK_SPEED_MAX_CHIP0_KEY__:          "pcie.link.speed.max.chip0",
		__PCIE_LINK_SPEED_MAX_CHIP1_KEY__:          "pcie.link.speed.max.chip1",
		__PCIE_LINK_SPEED_MAX_CHIP2_KEY__:          "pcie.link.speed.max.chip2",
		__PCIE_LINK_SPEED_CURRENT_CHIP0_KEY__:      "pcie.link.speed.current.chip0",
		__PCIE_LINK_SPEED_CURRENT_CHIP1_KEY__:      "pcie.link.speed.current.chip1",
		__PCIE_LINK_SPEED_CURRENT_CHIP2_KEY__:      "pcie.link.speed.current.chip2",
		__PCIE_LINK_GEN_MAX_CHIP0_KEY__:            "pcie.link.gen.max.chip0",
		__PCIE_LINK_GEN_MAX_CHIP1_KEY__:            "pcie.link.gen.max.chip1",
		__PCIE_LINK_GEN_MAX_CHIP2_KEY__:            "pcie.link.gen.max.chip2",
		__PCIE_LINK_GEN_CURRENT_CHIP0_KEY__:        "pcie.link.gen.current.chip0",
		__PCIE_LINK_GEN_CURRENT_CHIP1_KEY__:        "pcie.link.gen.current.chip1",
		__PCIE_LINK_GEN_CURRENT_CHIP2_KEY__:        "pcie.link.gen.current.chip2",
		__FAN_SPEED_KEY__:                          "fan.speed",
		__TEMPERATURE_CURRENT_CHIP0_KEY__:          "temperature.current.chip0",
		__TEMPERATURE_CURRENT_CHIP1_KEY__:          "temperature.current.chip1",
		__TEMPERATURE_CURRENT_CHIP2_KEY__:          "temperature.current.chip2",
		__VOLTAGE_CURRENT_CHIP0_KEY__:              "voltage.current.chip0",
		__VOLTAGE_CURRENT_CHIP1_KEY__:              "voltage.current.chip1",
		__VOLTAGE_CURRENT_CHIP2_KEY__:              "voltage.current.chip2",
		__VOLTAGE_BOARD_INPUT_KEY__:                "voltage.board.input",
		__CLOCKS_CURRENT_APU_CHIP0_KEY__:           "clocks.current.apu.chip0",
		__CLOCKS_CURRENT_APU_CHIP1_KEY__:           "clocks.current.apu.chip1",
		__CLOCKS_CURRENT_APU_CHIP2_KEY__:           "clocks.current.apu.chip2",
		__CLOCKS_CURRENT_CPU_CHIP0_KEY__:           "clocks.current.cpu.chip0",
		__CLOCKS_CURRENT_CPU_CHIP1_KEY__:           "clocks.current.cpu.chip1",
		__CLOCKS_CURRENT_CPU_CHIP2_KEY__:           "clocks.current.cpu.chip2",
		__CLOCKS_CURRENT_MEMORY_CHIP0_KEY__:        "clocks.current.memory.chip0",
		__CLOCKS_CURRENT_MEMORY_CHIP1_KEY__:        "clocks.current.memory.chip1",
		__CLOCKS_CURRENT_MEMORY_CHIP2_KEY__:        "clocks.current.memory.chip2",
		__CLOCKS_MAX_APU_CHIP0_KEY__:               "clocks.current.apu.max.chip0",
		__CLOCKS_MAX_APU_CHIP1_KEY__:               "clocks.current.apu.max.chip1",
		__CLOCKS_MAX_APU_CHIP2_KEY__:               "clocks.current.apu.max.chip2",
		__CLOCKS_MAX_CPU_CHIP0_KEY__:               "clocks.current.cpu.max.chip0",
		__CLOCKS_MAX_CPU_CHIP1_KEY__:               "clocks.current.cpu.max.chip1",
		__CLOCKS_MAX_CPU_CHIP2_KEY__:               "clocks.current.cpu.max.chip2",
		__CLOCKS_MAX_MEMORY_CHIP0_KEY__:            "clocks.current.memory.max.chip0",
		__CLOCKS_MAX_MEMORY_CHIP1_KEY__:            "clocks.current.memory.max.chip1",
		__CLOCKS_MAX_MEMORY_CHIP2_KEY__:            "clocks.current.memory.max.chip2",
		__POWER_DRAW__:                             "power.draw",
		__POWER_LIMIT__:                            "power.limit",
		__ECC_MODE_CURRENT_CHIP0_KEY__:             "ecc.mode.current.chip0",
		__ECC_MODE_CURRENT_CHIP1_KEY__:             "ecc.mode.current.chip1",
		__ECC_MODE_CURRENT_CHIP2_KEY__:             "ecc.mode.current.chip2",
		__ECC_ERRORS_CORRECTED_TOTAL_KEY__:         "ecc.errors.corrected.total",
		__ECC_ERRORS_CORRECTED_TOTAL_CHIP0_KEY__:   "ecc.errors.corrected.total.chip0",
		__ECC_ERRORS_CORRECTED_TOTAL_CHIP1_KEY__:   "ecc.errors.corrected.total.chip1",
		__ECC_ERRORS_CORRECTED_TOTAL_CHIP2_KEY__:   "ecc.errors.corrected.total.chip2",
		__ECC_ERRORS_UNCORRECTED_TOTAL_KEY__:       "ecc.errors.uncorrected.total",
		__ECC_ERRORS_UNCORRECTED_TOTAL_CHIP0_KEY__: "ecc.errors.uncorrected.total.chip0",
		__ECC_ERRORS_UNCORRECTED_TOTAL_CHIP1_KEY__: "ecc.errors.uncorrected.total.chip1",
		__ECC_ERRORS_UNCORRECTED_TOTAL_CHIP2_KEY__: "ecc.errors.uncorrected.total.chip2",
	}
)

const (
	__TIMESTAMP_KEY__                          = "timestamp"
	__BOARD_INDEX_KEY__                        = "board_index"
	__PRODUCT_NAME_KEY__                       = "name"
	__PRODUCT_NUMBER_KEY__                     = "product_number"
	__DRIVER_VERSION_KEY__                     = "driver_version"
	__FIRMWARE_VERSION_KEY__                   = "firmware_version"
	__SERIAL_NUMBER_KEY__                      = "serial_number"
	__CHIP_COUNT_KEY__                         = "chip_count"
	__CHIP_ID_LIST_KEY__                       = "chip_id_list"
	__CHIP_ID_CHIP_KEY__                       = "chip_id.chip"
	__CHIP_ID_CHIP0_KEY__                      = "chip_id.chip0"
	__CHIP_ID_CHIP1_KEY__                      = "chip_id.chip1"
	__CHIP_ID_CHIP2_KEY__                      = "chip_id.chip2"
	__UUID_LIST_KEY__                          = "uuid_list"
	__UUID_CHIP_KEY__                          = "uuid.chip"
	__UUID_CHIP0_KEY__                         = "uuid.chip0"
	__UUID_CHIP1_KEY__                         = "uuid.chip1"
	__UUID_CHIP2_KEY__                         = "uuid.chip2"
	__CHIP_INDEX_LIST_KEY__                    = "chip_index_list"
	__CHIP_INDEX_CHIP_KEY__                    = "chip_index.chip"
	__CHIP_INDEX_CHIP0_KEY__                   = "chip_index.chip0"
	__CHIP_INDEX_CHIP1_KEY__                   = "chip_index.chip1"
	__CHIP_INDEX_CHIP2_KEY__                   = "chip_index.chip2"
	__UTILIZATION_APU_TOTAL_KEY__              = "utilization.apu.total"
	__UTILIZATION_APU_CHIPS_KEY__              = "utilization.apu.chips"
	__UTILIZATION_APU_CHIP_KEY__               = "utilization.apu.chip"
	__UTILIZATION_APU_CHIP0_KEY__              = "utilization.apu.chip0"
	__UTILIZATION_APU_CHIP1_KEY__              = "utilization.apu.chip1"
	__UTILIZATION_APU_CHIP2_KEY__              = "utilization.apu.chip2"
	__UTILIZATION_CPU_TOTAL_KEY__              = "utilization.cpu.total"
	__UTILIZATION_CPU_CHIPS_KEY__              = "utilization.cpu.chips"
	__UTILIZATION_CPU_CHIP_KEY__               = "utilization.cpu.chip"
	__UTILIZATION_CPU_CHIP0_KEY__              = "utilization.cpu.chip0"
	__UTILIZATION_CPU_CHIP1_KEY__              = "utilization.cpu.chip1"
	__UTILIZATION_CPU_CHIP2_KEY__              = "utilization.cpu.chip2"
	__UTILIZATION_VIC_TOTAL_KEY__              = "utilization.vic.total"
	__UTILIZATION_VIC_CHIPS_KEY__              = "utilization.vic.chips"
	__UTILIZATION_VIC_CHIP_KEY__               = "utilization.vic.chip"
	__UTILIZATION_VIC_CHIP0_KEY__              = "utilization.vic.chip0"
	__UTILIZATION_VIC_CHIP1_KEY__              = "utilization.vic.chip1"
	__UTILIZATION_VIC_CHIP2_KEY__              = "utilization.vic.chip2"
	__UTILIZATION_MEMORY_TOTAL_KEY__           = "utilization.memory.total"
	__UTILIZATION_MEMORY_CHIPS_KEY__           = "utilization.memory.chips"
	__UTILIZATION_MEMORY_CHIP_KEY__            = "utilization.memory.chip"
	__UTILIZATION_MEMORY_CHIP0_KEY__           = "utilization.memory.chip0"
	__UTILIZATION_MEMORY_CHIP1_KEY__           = "utilization.memory.chip1"
	__UTILIZATION_MEMORY_CHIP2_KEY__           = "utilization.memory.chip2"
	__UTILIZATION_IPE_FPS_TOTAL_KEY__          = "utilization.ipeFps.total"
	__UTILIZATION_IPE_FPS_CHIPS_KEY__          = "utilization.ipeFps.chips"
	__UTILIZATION_IPE_FPS_CHIP_KEY__           = "utilization.ipeFps.chip"
	__UTILIZATION_IPE_FPS_CHIP0_KEY__          = "utilization.ipeFps.chip0"
	__UTILIZATION_IPE_FPS_CHIP1_KEY__          = "utilization.ipeFps.chip1"
	__UTILIZATION_IPE_FPS_CHIP2_KEY__          = "utilization.ipeFps.chip2"
	__PCI_CHIPS__KEY__                         = "pci.chips"
	__PCI_VENDOR_ID_KEY__                      = "pci.vendor_id"
	__PCI_VENDOR_ID_CHIP_KEY__                 = "pci.vendor_id.chip"
	__PCI_VENDOR_ID_CHIP0_KEY__                = "pci.vendor_id.chip0"
	__PCI_VENDOR_ID_CHIP1_KEY__                = "pci.vendor_id.chip1"
	__PCI_VENDOR_ID_CHIP2_KEY__                = "pci.vendor_id.chip2"
	__PCI_SUB_VENDOR_ID_KEY__                  = "pci.sub_vendor_id"
	__PCI_SUB_VENDOR_ID_CHIP_KEY__             = "pci.sub_vendor_id.chip"
	__PCI_SUB_VENDOR_ID_CHIP0_KEY__            = "pci.sub_vendor_id.chip0"
	__PCI_SUB_VENDOR_ID_CHIP1_KEY__            = "pci.sub_vendor_id.chip1"
	__PCI_SUB_VENDOR_ID_CHIP2_KEY__            = "pci.sub_vendor_id.chip2"
	__PCI_BUS_KEY__                            = "pci.bus"
	__PCI_BUS_CHIP_KEY__                       = "pci.bus.chip"
	__PCI_BUS_CHIP0_KEY__                      = "pci.bus.chip0"
	__PCI_BUS_CHIP1_KEY__                      = "pci.bus.chip1"
	__PCI_BUS_CHIP2_KEY__                      = "pci.bus.chip2"
	__PCI_DEVICE_ID_KEY__                      = "pci.device_id"
	__PCI_DEVICE_ID_CHIP_KEY__                 = "pci.device_id.chip"
	__PCI_DEVICE_ID_CHIP0_KEY__                = "pci.device_id.chip0"
	__PCI_DEVICE_ID_CHIP1_KEY__                = "pci.device_id.chip1"
	__PCI_DEVICE_ID_CHIP2_KEY__                = "pci.device_id.chip2"
	__PCI_SUB_DEVICE_ID_KEY__                  = "pci.sub_device_id"
	__PCI_SUB_DEVICE_ID_CHIP_KEY__             = "pci.sub_device_id.chip"
	__PCI_SUB_DEVICE_ID_CHIP0_KEY__            = "pci.sub_device_id.chip0"
	__PCI_SUB_DEVICE_ID_CHIP1_KEY__            = "pci.sub_device_id.chip1"
	__PCI_SUB_DEVICE_ID_CHIP2_KEY__            = "pci.sub_device_id.chip2"
	__PCI_DEVICE_KEY__                         = "pci.device"
	__PCI_DEVICE_CHIP_KEY__                    = "pci.device.chip"
	__PCI_DEVICE_CHIP0_KEY__                   = "pci.device.chip0"
	__PCI_DEVICE_CHIP1_KEY__                   = "pci.device.chip1"
	__PCI_DEVICE_CHIP2_KEY__                   = "pci.device.chip2"
	__PCI_SUB_FUNCTION_KEY__                   = "pci.function"
	__PCI_SUB_FUNCTION_CHIP_KEY__              = "pci.function.chip"
	__PCI_FUNCTION_CHIP0_KEY__                 = "pci.function.chip0"
	__PCI_FUNCTION_CHIP1_KEY__                 = "pci.function.chip1"
	__PCI_FUNCTION_CHIP2_KEY__                 = "pci.function.chip2"
	__PCI_NUMA_NODE_ID_KEY__                   = "pci.numa.node_id"
	__PCI_NUMA_NODE_ID_CHIP_KEY__              = "pci.numa.node_id.chip"
	__PCI_NUMA_NODE_ID_CHIP0_KEY__             = "pci.numa.node_id.chip0"
	__PCI_NUMA_NODE_ID_CHIP1_KEY__             = "pci.numa.node_id.chip1"
	__PCI_NUMA_NODE_ID_CHIP2_KEY__             = "pci.numa.node_id.chip2"
	__PCI_NUMA_CPU_KEY__                       = "pci.numa.cpu"
	__PCI_NUMA_CPU_CHIP_KEY__                  = "pci.numa.cpu.chip"
	__PCI_NUMA_CPU_CHIP0_KEY__                 = "pci.numa.cpu.chip0"
	__PCI_NUMA_CPU_CHIP1_KEY__                 = "pci.numa.cpu.chip1"
	__PCI_NUMA_CPU_CHIP2_KEY__                 = "pci.numa.cpu.chip2"
	__PCIE_LINK_SPEED_MAX_KEY__                = "pcie.link.speed.max"
	__PCIE_LINK_SPEED_MAX_CHIP_KEY__           = "pcie.link.speed.max.chip"
	__PCIE_LINK_SPEED_MAX_CHIP0_KEY__          = "pcie.link.speed.max.chip0"
	__PCIE_LINK_SPEED_MAX_CHIP1_KEY__          = "pcie.link.speed.max.chip1"
	__PCIE_LINK_SPEED_MAX_CHIP2_KEY__          = "pcie.link.speed.max.chip2"
	__PCIE_LINK_SPEED_CURRENT_KEY__            = "pcie.link.speed.current"
	__PCIE_LINK_SPEED_CURRENT_CHIP_KEY__       = "pcie.link.speed.current.chip"
	__PCIE_LINK_SPEED_CURRENT_CHIP0_KEY__      = "pcie.link.speed.current.chip0"
	__PCIE_LINK_SPEED_CURRENT_CHIP1_KEY__      = "pcie.link.speed.current.chip1"
	__PCIE_LINK_SPEED_CURRENT_CHIP2_KEY__      = "pcie.link.speed.current.chip2"
	__PCIE_LINK_GEN_MAX_KEY__                  = "pcie.link.gen.max"
	__PCIE_LINK_GEN_MAX_CHIP_KEY__             = "pcie.link.gen.max.chip"
	__PCIE_LINK_GEN_MAX_CHIP0_KEY__            = "pcie.link.gen.max.chip0"
	__PCIE_LINK_GEN_MAX_CHIP1_KEY__            = "pcie.link.gen.max.chip1"
	__PCIE_LINK_GEN_MAX_CHIP2_KEY__            = "pcie.link.gen.max.chip2"
	__PCIE_LINK_GEN_CURRENT_KEY__              = "pcie.link.gen.current"
	__PCIE_LINK_GEN_CURRENT_CHIP_KEY__         = "pcie.link.gen.current.chip"
	__PCIE_LINK_GEN_CURRENT_CHIP0_KEY__        = "pcie.link.gen.current.chip0"
	__PCIE_LINK_GEN_CURRENT_CHIP1_KEY__        = "pcie.link.gen.current.chip1"
	__PCIE_LINK_GEN_CURRENT_CHIP2_KEY__        = "pcie.link.gen.current.chip2"
	__FAN_SPEED_KEY__                          = "fan.speed"
	__TEMPERATURE_CURRENT_CHIPS_KEY__          = "temperature.current.chips"
	__TEMPERATURE_CURRENT_CHIP_KEY__           = "temperature.current.chip"
	__TEMPERATURE_CURRENT_CHIP0_KEY__          = "temperature.current.chip0"
	__TEMPERATURE_CURRENT_CHIP1_KEY__          = "temperature.current.chip1"
	__TEMPERATURE_CURRENT_CHIP2_KEY__          = "temperature.current.chip2"
	__VOLTAGE_CURRENT_CHIPS_KEY__              = "voltage.current.chips"
	__VOLTAGE_CURRENT_CHIP_KEY__               = "voltage.current.chip"
	__VOLTAGE_CURRENT_CHIP0_KEY__              = "voltage.current.chip0"
	__VOLTAGE_CURRENT_CHIP1_KEY__              = "voltage.current.chip1"
	__VOLTAGE_CURRENT_CHIP2_KEY__              = "voltage.current.chip2"
	__VOLTAGE_BOARD_INPUT_KEY__                = "voltage.board.input"
	__CLOCKS_KEY__                             = "clocks"
	__CLOCKS_CURRENT_APU_CHIPS_KEY__           = "clocks.current.apu.chips"
	__CLOCKS_CURRENT_APU_KEY__                 = "clocks.current.apu"
	__CLOCKS_CURRENT_APU_CHIP_KEY__            = "clocks.current.apu.chip"
	__CLOCKS_CURRENT_APU_CHIP0_KEY__           = "clocks.current.apu.chip0"
	__CLOCKS_CURRENT_APU_CHIP1_KEY__           = "clocks.current.apu.chip1"
	__CLOCKS_CURRENT_APU_CHIP2_KEY__           = "clocks.current.apu.chip2"
	__CLOCKS_CURRENT_CPU_CHIPS_KEY__           = "clocks.current.cpu.chips"
	__CLOCKS_CURRENT_CPU_KEY__                 = "clocks.current.cpu"
	__CLOCKS_CURRENT_CPU_CHIP_KEY__            = "clocks.current.cpu.chip"
	__CLOCKS_CURRENT_CPU_CHIP0_KEY__           = "clocks.current.cpu.chip0"
	__CLOCKS_CURRENT_CPU_CHIP1_KEY__           = "clocks.current.cpu.chip1"
	__CLOCKS_CURRENT_CPU_CHIP2_KEY__           = "clocks.current.cpu.chip2"
	__CLOCKS_CURRENT_MEMORY_CHIPS_KEY__        = "clocks.current.memory.chips"
	__CLOCKS_CURRENT_MEMORY_KEY__              = "clocks.current.memory"
	__CLOCKS_CURRENT_MEMORY_CHIP_KEY__         = "clocks.current.memory.chip"
	__CLOCKS_CURRENT_MEMORY_CHIP0_KEY__        = "clocks.current.memory.chip0"
	__CLOCKS_CURRENT_MEMORY_CHIP1_KEY__        = "clocks.current.memory.chip1"
	__CLOCKS_CURRENT_MEMORY_CHIP2_KEY__        = "clocks.current.memory.chip2"
	__CLOCKS_MAX_APU_CHIPS_KEY__               = "clocks.current.apu.max.chips"
	__CLOCKS_MAX_APU_CHIP_KEY__                = "clocks.current.apu.max.chip"
	__CLOCKS_MAX_APU_CHIP0_KEY__               = "clocks.current.apu.max.chip0"
	__CLOCKS_MAX_APU_CHIP1_KEY__               = "clocks.current.apu.max.chip1"
	__CLOCKS_MAX_APU_CHIP2_KEY__               = "clocks.current.apu.max.chip2"
	__CLOCKS_MAX_CPU_CHIPS_KEY__               = "clocks.current.cpu.max.chips"
	__CLOCKS_MAX_CPU_CHIP_KEY__                = "clocks.current.cpu.max.chip"
	__CLOCKS_MAX_CPU_CHIP0_KEY__               = "clocks.current.cpu.max.chip0"
	__CLOCKS_MAX_CPU_CHIP1_KEY__               = "clocks.current.cpu.max.chip1"
	__CLOCKS_MAX_CPU_CHIP2_KEY__               = "clocks.current.cpu.max.chip2"
	__CLOCKS_MAX_MEMORY_CHIPS_KEY__            = "clocks.current.memory.max.chips"
	__CLOCKS_MAX_MEMORY_CHIP_KEY__             = "clocks.current.memory.max.chip"
	__CLOCKS_MAX_MEMORY_CHIP0_KEY__            = "clocks.current.memory.max.chip0"
	__CLOCKS_MAX_MEMORY_CHIP1_KEY__            = "clocks.current.memory.max.chip1"
	__CLOCKS_MAX_MEMORY_CHIP2_KEY__            = "clocks.current.memory.max.chip2"
	__POWER_DRAW__                             = "power.draw"
	__POWER_LIMIT__                            = "power.limit"
	__ECC_MODE_CURRENT_CHIPS_KEY__             = "ecc.mode.current.chips"
	__ECC_MODE_CURRENT_CHIP_KEY__              = "ecc.mode.current.chip"
	__ECC_MODE_CURRENT_CHIP0_KEY__             = "ecc.mode.current.chip0"
	__ECC_MODE_CURRENT_CHIP1_KEY__             = "ecc.mode.current.chip1"
	__ECC_MODE_CURRENT_CHIP2_KEY__             = "ecc.mode.current.chip2"
	__ECC_ERRORS_CORRECTED_TOTAL_KEY__         = "ecc.errors.corrected.total"
	__ECC_ERRORS_CORRECTED_TOTAL_CHIPS_KEY__   = "ecc.errors.corrected.total.chips"
	__ECC_ERRORS_CORRECTED_TOTAL_CHIP_KEY__    = "ecc.errors.corrected.total.chip"
	__ECC_ERRORS_CORRECTED_TOTAL_CHIP0_KEY__   = "ecc.errors.corrected.total.chip0"
	__ECC_ERRORS_CORRECTED_TOTAL_CHIP1_KEY__   = "ecc.errors.corrected.total.chip1"
	__ECC_ERRORS_CORRECTED_TOTAL_CHIP2_KEY__   = "ecc.errors.corrected.total.chip2"
	__ECC_ERRORS_UNCORRECTED_TOTAL_KEY__       = "ecc.errors.uncorrected.total"
	__ECC_ERRORS_UNCORRECTED_TOTAL_CHIPS_KEY__ = "ecc.errors.uncorrected.total.chips"
	__ECC_ERRORS_UNCORRECTED_TOTAL_CHIP_KEY__  = "ecc.errors.uncorrected.total.chip"
	__ECC_ERRORS_UNCORRECTED_TOTAL_CHIP0_KEY__ = "ecc.errors.uncorrected.total.chip0"
	__ECC_ERRORS_UNCORRECTED_TOTAL_CHIP1_KEY__ = "ecc.errors.uncorrected.total.chip1"
	__ECC_ERRORS_UNCORRECTED_TOTAL_CHIP2_KEY__ = "ecc.errors.uncorrected.total.chip2"
)
