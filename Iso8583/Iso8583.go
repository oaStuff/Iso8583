package Iso8583
import "com.aihe/Iso8583/FieldValidator"

var(
	defaultTemplate TemplateDef
)

type Iso8583 struct  {
	AMessage
}

func NewIso8583() *Iso8583 {
	return &Iso8583{AMessage:*NewAMessage(NewTemplate(defaultTemplate))}
}

func NewIso8583WithTemplate(template TemplateDef) *Iso8583 {
	return &Iso8583{AMessage:*NewAMessage(NewTemplate(template))}
}


func init()  {
	defaultTemplate = TemplateDef{
		BIT_002_PAN 			:	AsciiVar(2, 19, fieldValidator.N()),
		BIT_003_PROC_CODE		:	AsciiFixed(6, fieldValidator.N()),
		BIT_004_TRAN_AMOUNT		:	AsciiNumeric(12),
		BIT_005_SETTLE_AMOUNT	:	AsciiNumeric(12),
		BIT_007_TRAN_DATE_TIME	:	AsciiFixed(10, fieldValidator.N()),
		BIT_009_CONVERSION_RATE_SETTLEMENT	:	AsciiFixed(8, fieldValidator.N()),
		BIT_011_SYS_TRACE_AUDIT_NUM			:	AsciiFixed(6, fieldValidator.N()),
		BIT_012_LOCAL_TRAN_TIME	:	AsciiFixed(6, fieldValidator.N()),
		BIT_013_LOCAL_TRAN_DATE	:	AsciiFixed(4, fieldValidator.N()),
		BIT_014_EXPIRATION_DATE	:	AsciiFixed(4, fieldValidator.N()),
		BIT_015_SELLTLEMENT_DATE:	AsciiFixed(4, fieldValidator.N()),
		BIT_016_CONVERSION_DATE	:	AsciiFixed(4, fieldValidator.N()),
		BIT_018_MERCHANT_TYPE	:	AsciiFixed(4, fieldValidator.N()),
		BIT_022_POS_ENTRY_MODE	: 	AsciiFixed(3, fieldValidator.N()),
		BIT_023_CARD_SEQUENCE_NUM			:	AsciiFixed(3, fieldValidator.N()),
		BIT_025_POS_CONDITION_CODE			:	AsciiFixed(2, fieldValidator.N()),
		BIT_026_POS_PIN_CAPTURE_CODE		:	AsciiFixed(2, fieldValidator.N()),
		BIT_027_AUTH_ID_RSP		:	AsciiFixed(1, fieldValidator.N()),
		BIT_028_TRAN_FEE_AMOUNT	:	AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		BIT_029_SETTLEMENT_FEE_AMOUNT		:	AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		BIT_030_TRAN_PROC_FEE_AMOUNT		:	AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		BIT_031_SETTLEMENT_PROC_FEE_AMOUNT	:	AsciiFixed(9, fieldValidator.Rev87AmountValidator()),
		BIT_032_ACQUIRING_INST_ID_CODE		:	AsciiVar(2, 11, fieldValidator.N()),
		BIT_033_FORWARDING_INT_ID_CODE		:	AsciiVar(2, 11, fieldValidator.N()),
		BIT_035_TRACK_2_DATA	:	AsciiVar(2, 37, fieldValidator.Ans()),
		BIT_037_RETRIEVAL_REF_NUM			:	AsciiFixed(12, fieldValidator.N()),
		BIT_038_AUTH_ID_RESPONSE			:	AsciiFixed(6, fieldValidator.N()),
		BIT_039_RESPONSE_CODE	: AsciiFixed(2, fieldValidator.An()),
		BIT_040_SERVICE_RESTRICTION_CODE	:	AsciiFixed(3, fieldValidator.N()),
		BIT_041_CARD_ACCEPTOR_TERMINAL_ID	:	AsciiFixed(8, fieldValidator.Ans()),
		BIT_042_CARD_ACCEPTOR_ID_CODE		:	AsciiFixed(15, fieldValidator.Ans()),
		BIT_043_CARD_ACCEPTOR_NAME_LOCATION	:	AsciiFixed(40, fieldValidator.Ans()),
		BIT_044_ADDITIONAL_RESPONSE_DATA	:	AsciiVar(2,25,fieldValidator.Ans()),
		BIT_045_TRACK_1_DATA	:	AsciiVar(2,76,fieldValidator.Ans()),
		BIT_048_ADDITIONAL_DATA	:	AsciiVar(3,999,fieldValidator.Ans()),
		BIT_049_TRAN_CURRENCY_CODE			:	AsciiFixed(3, fieldValidator.N()),
		BIT_050_SETTLEMENT_CURRENCY_CODE	:	AsciiFixed(3, fieldValidator.N()),
		BIT_052_PIN_DATA		:	BinaryFixed(8),
		BIT_053_SECURITY_RELATED_CONTROL_INFORMATION	:	BinaryFixed(48),
		BIT_054_ADDITIONAL_AMOUNTS			:	AsciiVar(3,120,fieldValidator.An()),
		BIT_056_MESSAGE_REASON_CODE			:	AsciiVar(3,4,fieldValidator.N()),
		BIT_057_AUTHORISATION_LIFE_CYCLE	:	AsciiVar(3,3,fieldValidator.N()),
		BIT_058_AUTHORISING_AGENT_INSTITUTION			:	AsciiVar(3,11,fieldValidator.Anp()),
		BIT_066_SETTLEMENT_CODE				:	AsciiFixed(1, fieldValidator.N()),
		BIT_067_EXTENDED_PAYMENT_CODE		:	AsciiFixed(2, fieldValidator.N()),
		BIT_070_NETWORK_MANAGEMENT_INFORMATION_CODE		:	AsciiFixed(3, fieldValidator.N()),
		BIT_073_DATE_ACTION		:	AsciiFixed(6, fieldValidator.N()),
		BIT_074_CREDITS_NUMBER	:	AsciiFixed(10, fieldValidator.N()),
		BIT_075_CREDITS_REVERSAL_NUMBER		:	AsciiFixed(10, fieldValidator.N()),
		BIT_076_DEBITS_NUMBER	:	AsciiFixed(10, fieldValidator.N()),
		BIT_077_DEBITS_REVERSAL_NUMBER		:	AsciiFixed(10, fieldValidator.N()),
		BIT_078_TRANSFER_NUMBER	:	AsciiFixed(10, fieldValidator.N()),
		BIT_079_TRANSFER_REVERSAL_NUMBER	:	AsciiFixed(10, fieldValidator.N()),
		BIT_080_INQUIRIES_NUMBER			:	AsciiFixed(10, fieldValidator.N()),
		BIT_081_AUTHORISATIONS_NUMBER		:	AsciiFixed(10, fieldValidator.N()),
		BIT_082_CREDITS_PROCESSING_FEE_AMOUNT			:	AsciiFixed(12, fieldValidator.N()),
		BIT_083_CREDITS_TRANSACTION_FEE_AMOUNT			:	AsciiFixed(12, fieldValidator.N()),
		BIT_084_DEBITS_PROCESSING_FEE_AMOUNT			:	AsciiFixed(12, fieldValidator.N()),
		BIT_085_DEBITS_TRANSACTION_FEE_AMOUNT			:	AsciiFixed(12, fieldValidator.N()),
		BIT_086_CREDITS_AMOUNT	:	AsciiFixed(16, fieldValidator.N()),
		BIT_087_CREDITS_REVERSAL_AMOUNT		:	AsciiFixed(16, fieldValidator.N()),
		BIT_088_DEBITS_AMOUNT	:	AsciiFixed(16, fieldValidator.N()),
		BIT_089_DEBITS_REVERSAL_AMOUNT		:	AsciiFixed(16, fieldValidator.N()),
		BIT_090_ORIGINAL_DATA_ELEMENTS		:	AsciiFixed(42, fieldValidator.N()),
		BIT_091_FILE_UPDATE_CODE			:	AsciiFixed(1, fieldValidator.An()),
		BIT_095_REPLACEMENT_AMOUNTS			:	AsciiFixed(42, fieldValidator.Ans()),
		BIT_097_AMOUNT_NET_SETTLEMENT		:	AsciiFixed(17, fieldValidator.Rev87AmountValidator()),
		BIT_098_PAYEE			:	AsciiFixed(25, fieldValidator.Ans()),
		BIT_100_RECEIVING_INST_ID_CODE		:	AsciiVar(2,11,fieldValidator.Ans()),
		BIT_101_FILE_NAME		:	AsciiVar(2,17,fieldValidator.Ans()),
		BIT_102_ACCOUNT_ID_1	:	AsciiVar(2,28,fieldValidator.Ans()),
		BIT_103_ACCOUNT_ID_2	:	AsciiVar(2,28,fieldValidator.Ans()),
		BIT_118_PAYMENTS_NUMBER	:	AsciiVar(3,30,fieldValidator.N()),
		BIT_119_PAYMENTS_REVERSAL_NUMBER	:	AsciiVar(3,10,fieldValidator.N()),
	}
}