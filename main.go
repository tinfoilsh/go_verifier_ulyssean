package main

import (
	"log"

	"github.com/tinfoilsh/verifier/attestation"
)

var knownMeasurement = &attestation.Measurement{
	Type:      attestation.SevGuestV1,
	Registers: []string{"a1a0b8889c305aa0b12899b2565a032287da44ec996a06d03c835a200684d4c4edadfa67997374fba2270a79cfb013a2"},
}

func main() {
	enclaveAttestation, err := attestation.Fetch("ulyssean-demo.delta.tinfoil.sh")
	if err != nil {
		log.Fatalf("failed to fetch enclave measurements: %v", err)
	}
	verification, err := enclaveAttestation.Verify()
	if err != nil {
		log.Fatalf("failed to verify enclave measurements: %v", err)
	}

	if err := verification.Measurement.Equals(knownMeasurement); err != nil {
		log.Fatalf("measurement mismatch: %v", err)
	}

	log.Printf("Measurement match: %+v", verification.Measurement)
}
