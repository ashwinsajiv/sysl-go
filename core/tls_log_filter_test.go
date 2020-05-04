package core

// func TestTLSLogFilter_Write(t *testing.T) {
// 	type testData struct {
// 		in    string
// 		level logrus.Level
// 	}

// 	for i, tt := range []testData{
// 		{
// 			in:    "hit hit\n",
// 			level: logrus.DebugLevel,
// 		},
// 		{
// 			in:    "misssssss\n",
// 			level: logrus.WarnLevel,
// 		},
// 	} {
// 		tt := tt
// 		t.Run(fmt.Sprintf("TestTLSLogFilter_Write-%d", i), func(t *testing.T) {
// 			ctx, hook := common.NewTestContextWithLoggerHook()

// 			re := regexp.MustCompile(`hit`)
// 			writer := &TLSLogFilter{logger, re}
// 			serverLogger := log.New(writer, "", 0)

// 			serverLogger.Printf(tt.in)

// 			require.Equal(t, 1, len(hook.Entries))
// 		})
// 	}
// }
