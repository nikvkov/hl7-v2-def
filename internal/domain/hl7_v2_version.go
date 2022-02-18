package domain

type Hl7V2Version string

const (
	Hl7v2version21  Hl7V2Version = "HL7v2.1"
	HL7v2version2                = "HL7v2.2"
	HL7v2version23               = "HL7v2.3"
	HL7v2version231              = "HL7v2.3.1"
	HL7v2version24               = "HL7v2.4"
	HL7v2version25               = "HL7v2.5"
	HL7v2version251              = "HL7v2.5.1"
	HL7v2version26               = "HL7v2.6"
	HL7v2version27               = "HL7v2.7"
	HL7v2version271              = "HL7v2.7.1"
	HL7v2version28               = "HL7v2.8"
)

func (v *Hl7V2Version) String() string {
	return string(*v)
}
