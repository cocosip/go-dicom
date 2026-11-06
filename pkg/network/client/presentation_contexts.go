// Copyright (c) 2026 go-dicom contributors.
// Licensed under the MIT License.

package client

import "github.com/cocosip/go-dicom/pkg/dicom/uid"

// PresentationContextSpec describes one requested presentation context.
//
// It is the value-object form of a presentation context request and is shared by
// Client (AddPresentationContexts) and ManagedClient jobs, so contexts assembled
// once can be reused by either client style.
type PresentationContextSpec struct {
	AbstractSyntax   string
	TransferSyntaxes []string
	SCURole          bool
	SCPRole          bool
}

// NewPresentationContextSpec builds one presentation context request. When either
// scuRole or scpRole is true an SCP/SCU Role Selection sub-item is proposed.
func NewPresentationContextSpec(abstractSyntax string, scuRole, scpRole bool, transferSyntaxes ...string) PresentationContextSpec {
	return PresentationContextSpec{
		AbstractSyntax:   abstractSyntax,
		TransferSyntaxes: append([]string(nil), transferSyntaxes...),
		SCURole:          scuRole,
		SCPRole:          scpRole,
	}
}

// StoragePresentationContextSpecs returns SCP-role presentation context specs for
// every active standard Storage SOP Class in category. It is the Go equivalent of
// fo-dicom's DicomPresentationContext.GetScpRolePresentationContextsFromStorageUids.
func StoragePresentationContextSpecs(category uid.StorageCategory, transferSyntaxes ...string) []PresentationContextSpec {
	return StoragePresentationContextSpecsFor(uid.StorageSOPClassUIDs(category), transferSyntaxes...)
}

// StoragePresentationContextSpecsFor returns SCP-role presentation context specs
// for the given Storage SOP Class UIDs. It is the Go equivalent of fo-dicom's
// DicomPresentationContext.GetScpRolePresentationContext.
func StoragePresentationContextSpecsFor(sopClassUIDs []string, transferSyntaxes ...string) []PresentationContextSpec {
	specs := make([]PresentationContextSpec, 0, len(sopClassUIDs))
	for _, sopClassUID := range sopClassUIDs {
		specs = append(specs, NewPresentationContextSpec(sopClassUID, false, true, transferSyntaxes...))
	}
	return specs
}

// AddPresentationContexts applies pre-built presentation context specs. It lets a
// caller assemble arbitrary contexts (including private SOP Classes and explicit
// SCU/SCP roles) using the same value type that ManagedClient jobs consume.
func (c *Client) AddPresentationContexts(specs []PresentationContextSpec) error {
	for _, spec := range specs {
		var err error
		if spec.SCURole || spec.SCPRole {
			err = c.AddPresentationContextWithRoles(spec.AbstractSyntax, spec.SCURole, spec.SCPRole, spec.TransferSyntaxes...)
		} else {
			err = c.AddPresentationContext(spec.AbstractSyntax, spec.TransferSyntaxes...)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
