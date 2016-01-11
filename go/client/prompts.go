// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"github.com/keybase/client/go/libkb"
)

const (
	PromptDescriptorGeneric libkb.PromptDescriptor = iota
	PromptDescriptorRevokePaperKeys
	PromptDescriptorReregister
	PromptDescriptorInviteOK
	PromptDescriptorPGPGenPushSecret
	PromptDescriptorDoctorWhichAccount
	PromptDescriptorDoctorSignOK
	PromptDescriptorGPGOKToAdd
	PromptDescriptorGPGConfirmDuplicateKey
	PromptDescriptorTrackAction
	PromptDescriptorTrackPublic
	PromptDescriptorProvePreWarning
	PromptDescriptorProveOKToCheck
	PromptDescriptorProveOverwriteOK
	PromptDescriptorLocksmithDeviceName
	PromptDescriptorLocksmithSigningOption
	PromptDescriptorGPGSelectKey
	PromptDescriptorLoginUsername
	PromptDescriptorLoginWritePaper
	PromptDescriptorLoginWallet
	PromptDescriptorSignupFullName
	PromptDescriptorSignupNotes
	PromptDescriptorSignupUsername
	PromptDescriptorSignupEmail
	PromptDescriptorSignupReenterPassphrase
	PromptDescriptorSignupDevice
	PromptDescriptorSignupCode
	PromptDescriptorChooseProvisioningMethod
	PromptDescriptorChooseGPGMethod
	PromptDescriptorChooseDeviceType
	PromptDescriptorProvisionPhrase
	PromptDescriptorProvisionDeviceName
	PromptDescriptorExportSecretKeyFromGPG
	PromptDescriptorDeprovisionWhichUser
	PromptDescriptorUpdateDo
	PromptDescriptorUpdateAuto
	PromptDescriptorUpdateSnooze
	PromptDescriptorDecryptInteractive
	PromptDescriptorPGPGenEnterID
)
