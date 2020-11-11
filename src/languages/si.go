package languages

func GetPackSI() Language {
	return Language{
		Iso639Code: "si",
		ConsoleOutputMessages: ConsoleOutputMessages{
			CantWriteFile:     "ගොනුව ලිවීමේ දෝෂයකි",
			FileAlreadyExists: "ගොනුව දැනටමත් පවතී.\nගොනුව නැවත ලිවීමට \"--force\" flag භාවිතා කරන්න",
			OverWritingFIle:   "\"--force\" flag බාවිතා කිරීම නිසා ගොනුව ආපසු ලියනු ලැබේ",
			FileWritten:       "ගොනුව සාර්ථකව ලියා ඇත",
		},
	}
}
