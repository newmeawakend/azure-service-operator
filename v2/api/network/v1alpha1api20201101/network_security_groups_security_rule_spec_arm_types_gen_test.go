// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_NetworkSecurityGroups_SecurityRule_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_SecurityRule_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_SpecARM, NetworkSecurityGroups_SecurityRule_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_SpecARM runs a test to see if a specific instance of NetworkSecurityGroups_SecurityRule_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroups_SecurityRule_SpecARM(subject NetworkSecurityGroups_SecurityRule_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_SecurityRule_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NetworkSecurityGroups_SecurityRule_SpecARM instances for property testing - lazily instantiated by
// NetworkSecurityGroups_SecurityRule_SpecARMGenerator()
var networkSecurityGroups_SecurityRule_SpecARMGenerator gopter.Gen

// NetworkSecurityGroups_SecurityRule_SpecARMGenerator returns a generator of NetworkSecurityGroups_SecurityRule_SpecARM instances for property testing.
// We first initialize networkSecurityGroups_SecurityRule_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroups_SecurityRule_SpecARMGenerator() gopter.Gen {
	if networkSecurityGroups_SecurityRule_SpecARMGenerator != nil {
		return networkSecurityGroups_SecurityRule_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_SpecARM(generators)
	networkSecurityGroups_SecurityRule_SpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_SpecARM(generators)
	networkSecurityGroups_SecurityRule_SpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SecurityRule_SpecARM{}), generators)

	return networkSecurityGroups_SecurityRule_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroups_SecurityRule_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecurityRulePropertiesFormatARMGenerator())
}

func Test_SecurityRulePropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRulePropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRulePropertiesFormatARM, SecurityRulePropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRulePropertiesFormatARM runs a test to see if a specific instance of SecurityRulePropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRulePropertiesFormatARM(subject SecurityRulePropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRulePropertiesFormatARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SecurityRulePropertiesFormatARM instances for property testing - lazily instantiated by
// SecurityRulePropertiesFormatARMGenerator()
var securityRulePropertiesFormatARMGenerator gopter.Gen

// SecurityRulePropertiesFormatARMGenerator returns a generator of SecurityRulePropertiesFormatARM instances for property testing.
// We first initialize securityRulePropertiesFormatARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRulePropertiesFormatARMGenerator() gopter.Gen {
	if securityRulePropertiesFormatARMGenerator != nil {
		return securityRulePropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM(generators)
	securityRulePropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormatARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM(generators)
	AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormatARM(generators)
	securityRulePropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(SecurityRulePropertiesFormatARM{}), generators)

	return securityRulePropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRulePropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRulePropertiesFormat_Access_Allow, SecurityRulePropertiesFormat_Access_Deny))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRulePropertiesFormat_Direction_Inbound, SecurityRulePropertiesFormat_Direction_Outbound))
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_Ah,
		SecurityRulePropertiesFormat_Protocol_Esp,
		SecurityRulePropertiesFormat_Protocol_Icmp,
		SecurityRulePropertiesFormat_Protocol_Star,
		SecurityRulePropertiesFormat_Protocol_Tcp,
		SecurityRulePropertiesFormat_Protocol_Udp))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRulePropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(SubResourceARMGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(SubResourceARMGenerator())
}