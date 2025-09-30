// Package census simulates a simple system for collecting census data.
package census

// Resident holds information about a city resident.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string // expects keys like "street", "city", etc.
}

// NewResident creates and returns a new Resident.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

// HasRequiredInfo reports whether the resident has the required census fields filled in.
func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && r.Address != nil && r.Address["street"] != ""
}

// Delete clears all information from the resident record.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count returns the number of residents that have the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, r := range residents {
		if r.HasRequiredInfo() {
			count++
		}
	}
	return count
}
