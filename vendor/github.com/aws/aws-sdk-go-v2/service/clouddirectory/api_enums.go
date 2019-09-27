// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

type BatchReadExceptionType string

// Enum values for BatchReadExceptionType
const (
	BatchReadExceptionTypeValidationException             BatchReadExceptionType = "ValidationException"
	BatchReadExceptionTypeInvalidArnException             BatchReadExceptionType = "InvalidArnException"
	BatchReadExceptionTypeResourceNotFoundException       BatchReadExceptionType = "ResourceNotFoundException"
	BatchReadExceptionTypeInvalidNextTokenException       BatchReadExceptionType = "InvalidNextTokenException"
	BatchReadExceptionTypeAccessDeniedException           BatchReadExceptionType = "AccessDeniedException"
	BatchReadExceptionTypeNotNodeException                BatchReadExceptionType = "NotNodeException"
	BatchReadExceptionTypeFacetValidationException        BatchReadExceptionType = "FacetValidationException"
	BatchReadExceptionTypeCannotListParentOfRootException BatchReadExceptionType = "CannotListParentOfRootException"
	BatchReadExceptionTypeNotIndexException               BatchReadExceptionType = "NotIndexException"
	BatchReadExceptionTypeNotPolicyException              BatchReadExceptionType = "NotPolicyException"
	BatchReadExceptionTypeDirectoryNotEnabledException    BatchReadExceptionType = "DirectoryNotEnabledException"
	BatchReadExceptionTypeLimitExceededException          BatchReadExceptionType = "LimitExceededException"
	BatchReadExceptionTypeInternalServiceException        BatchReadExceptionType = "InternalServiceException"
)

func (enum BatchReadExceptionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BatchReadExceptionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BatchWriteExceptionType string

// Enum values for BatchWriteExceptionType
const (
	BatchWriteExceptionTypeInternalServiceException         BatchWriteExceptionType = "InternalServiceException"
	BatchWriteExceptionTypeValidationException              BatchWriteExceptionType = "ValidationException"
	BatchWriteExceptionTypeInvalidArnException              BatchWriteExceptionType = "InvalidArnException"
	BatchWriteExceptionTypeLinkNameAlreadyInUseException    BatchWriteExceptionType = "LinkNameAlreadyInUseException"
	BatchWriteExceptionTypeStillContainsLinksException      BatchWriteExceptionType = "StillContainsLinksException"
	BatchWriteExceptionTypeFacetValidationException         BatchWriteExceptionType = "FacetValidationException"
	BatchWriteExceptionTypeObjectNotDetachedException       BatchWriteExceptionType = "ObjectNotDetachedException"
	BatchWriteExceptionTypeResourceNotFoundException        BatchWriteExceptionType = "ResourceNotFoundException"
	BatchWriteExceptionTypeAccessDeniedException            BatchWriteExceptionType = "AccessDeniedException"
	BatchWriteExceptionTypeInvalidAttachmentException       BatchWriteExceptionType = "InvalidAttachmentException"
	BatchWriteExceptionTypeNotIndexException                BatchWriteExceptionType = "NotIndexException"
	BatchWriteExceptionTypeNotNodeException                 BatchWriteExceptionType = "NotNodeException"
	BatchWriteExceptionTypeIndexedAttributeMissingException BatchWriteExceptionType = "IndexedAttributeMissingException"
	BatchWriteExceptionTypeObjectAlreadyDetachedException   BatchWriteExceptionType = "ObjectAlreadyDetachedException"
	BatchWriteExceptionTypeNotPolicyException               BatchWriteExceptionType = "NotPolicyException"
	BatchWriteExceptionTypeDirectoryNotEnabledException     BatchWriteExceptionType = "DirectoryNotEnabledException"
	BatchWriteExceptionTypeLimitExceededException           BatchWriteExceptionType = "LimitExceededException"
	BatchWriteExceptionTypeUnsupportedIndexTypeException    BatchWriteExceptionType = "UnsupportedIndexTypeException"
)

func (enum BatchWriteExceptionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BatchWriteExceptionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConsistencyLevel string

// Enum values for ConsistencyLevel
const (
	ConsistencyLevelSerializable ConsistencyLevel = "SERIALIZABLE"
	ConsistencyLevelEventual     ConsistencyLevel = "EVENTUAL"
)

func (enum ConsistencyLevel) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConsistencyLevel) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DirectoryState string

// Enum values for DirectoryState
const (
	DirectoryStateEnabled  DirectoryState = "ENABLED"
	DirectoryStateDisabled DirectoryState = "DISABLED"
	DirectoryStateDeleted  DirectoryState = "DELETED"
)

func (enum DirectoryState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DirectoryState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FacetAttributeType string

// Enum values for FacetAttributeType
const (
	FacetAttributeTypeString   FacetAttributeType = "STRING"
	FacetAttributeTypeBinary   FacetAttributeType = "BINARY"
	FacetAttributeTypeBoolean  FacetAttributeType = "BOOLEAN"
	FacetAttributeTypeNumber   FacetAttributeType = "NUMBER"
	FacetAttributeTypeDatetime FacetAttributeType = "DATETIME"
	FacetAttributeTypeVariant  FacetAttributeType = "VARIANT"
)

func (enum FacetAttributeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FacetAttributeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FacetStyle string

// Enum values for FacetStyle
const (
	FacetStyleStatic  FacetStyle = "STATIC"
	FacetStyleDynamic FacetStyle = "DYNAMIC"
)

func (enum FacetStyle) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FacetStyle) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ObjectType string

// Enum values for ObjectType
const (
	ObjectTypeNode     ObjectType = "NODE"
	ObjectTypeLeafNode ObjectType = "LEAF_NODE"
	ObjectTypePolicy   ObjectType = "POLICY"
	ObjectTypeIndex    ObjectType = "INDEX"
)

func (enum ObjectType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ObjectType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RangeMode string

// Enum values for RangeMode
const (
	RangeModeFirst                   RangeMode = "FIRST"
	RangeModeLast                    RangeMode = "LAST"
	RangeModeLastBeforeMissingValues RangeMode = "LAST_BEFORE_MISSING_VALUES"
	RangeModeInclusive               RangeMode = "INCLUSIVE"
	RangeModeExclusive               RangeMode = "EXCLUSIVE"
)

func (enum RangeMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RangeMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RequiredAttributeBehavior string

// Enum values for RequiredAttributeBehavior
const (
	RequiredAttributeBehaviorRequiredAlways RequiredAttributeBehavior = "REQUIRED_ALWAYS"
	RequiredAttributeBehaviorNotRequired    RequiredAttributeBehavior = "NOT_REQUIRED"
)

func (enum RequiredAttributeBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RequiredAttributeBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RuleType string

// Enum values for RuleType
const (
	RuleTypeBinaryLength     RuleType = "BINARY_LENGTH"
	RuleTypeNumberComparison RuleType = "NUMBER_COMPARISON"
	RuleTypeStringFromSet    RuleType = "STRING_FROM_SET"
	RuleTypeStringLength     RuleType = "STRING_LENGTH"
)

func (enum RuleType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RuleType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateActionType string

// Enum values for UpdateActionType
const (
	UpdateActionTypeCreateOrUpdate UpdateActionType = "CREATE_OR_UPDATE"
	UpdateActionTypeDelete         UpdateActionType = "DELETE"
)

func (enum UpdateActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}