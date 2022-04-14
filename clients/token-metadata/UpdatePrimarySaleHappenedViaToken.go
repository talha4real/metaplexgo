// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Allows updating the primary sale boolean on Metadata solely through owning an account
// containing a token from the metadata's mint and being a signer on this transaction.
// A sort of limited authority for limited update capability that is required for things like
// Metaplex to work without needing full authority passing.
type UpdatePrimarySaleHappenedViaToken struct {

	// [0] = [WRITE] metadataKeyPDA
	// ··········· Metadata key (pda of ['metadata', program id, mint id])
	//
	// [1] = [SIGNER] owner
	// ··········· Owner on the token account
	//
	// [2] = [] container
	// ··········· Account containing tokens from the metadata's mint
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePrimarySaleHappenedViaTokenInstructionBuilder creates a new `UpdatePrimarySaleHappenedViaToken` instruction builder.
func NewUpdatePrimarySaleHappenedViaTokenInstructionBuilder() *UpdatePrimarySaleHappenedViaToken {
	nd := &UpdatePrimarySaleHappenedViaToken{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetMetadataKeyPDAAccount sets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *UpdatePrimarySaleHappenedViaToken) SetMetadataKeyPDAAccount(metadataKeyPDA ag_solanago.PublicKey) *UpdatePrimarySaleHappenedViaToken {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadataKeyPDA).WRITE()
	return inst
}

// GetMetadataKeyPDAAccount gets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *UpdatePrimarySaleHappenedViaToken) GetMetadataKeyPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
// Owner on the token account
func (inst *UpdatePrimarySaleHappenedViaToken) SetOwnerAccount(owner ag_solanago.PublicKey) *UpdatePrimarySaleHappenedViaToken {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// Owner on the token account
func (inst *UpdatePrimarySaleHappenedViaToken) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetContainerAccount sets the "container" account.
// Account containing tokens from the metadata's mint
func (inst *UpdatePrimarySaleHappenedViaToken) SetContainerAccount(container ag_solanago.PublicKey) *UpdatePrimarySaleHappenedViaToken {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(container)
	return inst
}

// GetContainerAccount gets the "container" account.
// Account containing tokens from the metadata's mint
func (inst *UpdatePrimarySaleHappenedViaToken) GetContainerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdatePrimarySaleHappenedViaToken) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdatePrimarySaleHappenedViaToken),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePrimarySaleHappenedViaToken) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePrimarySaleHappenedViaToken) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MetadataKeyPDA is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Container is not set")
		}
	}
	return nil
}

func (inst *UpdatePrimarySaleHappenedViaToken) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePrimarySaleHappenedViaToken")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("metadataKeyPDA", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     container", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdatePrimarySaleHappenedViaToken) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UpdatePrimarySaleHappenedViaToken) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUpdatePrimarySaleHappenedViaTokenInstruction declares a new UpdatePrimarySaleHappenedViaToken instruction with the provided parameters and accounts.
func NewUpdatePrimarySaleHappenedViaTokenInstruction(
	// Accounts:
	metadataKeyPDA ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	container ag_solanago.PublicKey) *UpdatePrimarySaleHappenedViaToken {
	return NewUpdatePrimarySaleHappenedViaTokenInstructionBuilder().
		SetMetadataKeyPDAAccount(metadataKeyPDA).
		SetOwnerAccount(owner).
		SetContainerAccount(container)
}
