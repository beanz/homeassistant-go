package types

type DeviceClass string

const (
	DeviceClassAqi                      DeviceClass = "aqi"
	DeviceClassBattery                  DeviceClass = "battery"
	DeviceClassCarbonDioxide            DeviceClass = "carbon_dioxide"
	DeviceClassCarbonMonoxide           DeviceClass = "carbon_monoxide"
	DeviceClassCurrent                  DeviceClass = "current"
	DeviceClassDate                     DeviceClass = "date"
	DeviceClassEnergy                   DeviceClass = "energy"
	DeviceClassGas                      DeviceClass = "gas"
	DeviceClassHumidity                 DeviceClass = "humidity"
	DeviceClassIlluminance              DeviceClass = "illuminance"
	DeviceClassMonetary                 DeviceClass = "monetary"
	DeviceClassNitrogenDioxide          DeviceClass = "nitrogen_dioxide"
	DeviceClassNitrogenMonoxide         DeviceClass = "nitrogen_monoxide"
	DeviceClassNitrousOxide             DeviceClass = "nitrous_oxide"
	DeviceClassOzone                    DeviceClass = "ozone"
	DeviceClassPm1                      DeviceClass = "pm1"
	DeviceClassPm10                     DeviceClass = "pm10"
	DeviceClassPm25                     DeviceClass = "pm25"
	DeviceClassPowerFactor              DeviceClass = "power_factor"
	DeviceClassPower                    DeviceClass = "power"
	DeviceClassPressure                 DeviceClass = "pressure"
	DeviceClassSignalStrength           DeviceClass = "signal_strength"
	DeviceClassSulphurDioxide           DeviceClass = "sulphur_dioxide"
	DeviceClassTemperature              DeviceClass = "temperature"
	DeviceClassTimestamp                DeviceClass = "timestamp"
	DeviceClassVolatileOrganicCompounds DeviceClass = "volatile_organic_compounds"
	DeviceClassVoltage                  DeviceClass = "voltage"
)
