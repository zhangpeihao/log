package log

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	fmt.Println("================= TestLog ====================")
	counters := []string{"a", "b"}
	logger := NewLogger(".", "test", counters, 2, 10, true)
	// Default log level is Fatal
	logger.Debug("!!!!! logger.Debug(): no DEBUG log output in default level\n")
	logger.Debugf("!!!!! logger.Debugf(): no DEBUG log output in default level\n")
	logger.Debugln("!!!!! logger.Debugln(): no DEBUG log output in default level")

	logger.Print("!!!!! logger.Print(): no Print log output in default level\n")
	logger.Printf("!!!!! logger.Printf(): no Print log output in default level\n")
	logger.Println("!!!!! logger.Println(): no Print log output in default level")

	logger.Warning("!!!!! logger.Waining(): no WARNING log output in default level\n")
	logger.Warningf("!!!!! logger.Wainingf(): no WARNING log output in default level\n")
	logger.Warningln("!!!!! logger.Wainingln(): no WARNING log output in default level")

	logger.Fatal("1. logger.Fatal(): FATAL log output in default level\n")
	logger.Fatalf("2. logger.Fatalf(): FATAL log output in default level\n")
	logger.Fatalln("3. logger.Fatalln(): FATAL log output in default level")

	// Set level to WARNING
	logger.SetMainLevel(LOG_LEVEL_WARNING)
	logger.Debug("!!!!! logger.Debug(): no DEBUG log output in WARNING level\n")
	logger.Debugf("!!!!! logger.Debugf(): no DEBUG log output in WARNING level\n")
	logger.Debugln("!!!!! logger.Debugln(): no DEBUG log output in WARNING level")

	logger.Print("!!!!! logger.Print(): no Print log output in WARNING level\n")
	logger.Printf("!!!!! logger.Printf(): no Print log output in WARNING level\n")
	logger.Println("!!!!! logger.Println(): no Print log output in WARNING level")

	logger.Warning("4. logger.Waining(): WARNING log output in WARNING level\n")
	logger.Warningf("5. logger.Wainingf(): WARNING log output in WARNING level\n")
	logger.Warningln("6. logger.Wainingln(): WARNING log output in WARNING level")

	logger.Fatal("7. logger.Fatal(): FATAL log output in WARNING level\n")
	logger.Fatalf("8. logger.Fatalf(): FATAL log output in WARNING level\n")
	logger.Fatalln("9. logger.Fatalln(): FATAL log output in WARNING level")

	// Set level to TRACE
	logger.SetMainLevel(LOG_LEVEL_TRACE)
	logger.Debug("!!!!! logger.Debug(): no DEBUG log output in TRACE level\n")
	logger.Debugf("!!!!! logger.Debugf(): no DEBUG log output in TRACE level\n")
	logger.Debugln("!!!!! logger.Debugln(): no DEBUG log output in TRACE level")

	logger.Print("10. logger.Print(): Print log output in TRACE level\n")
	logger.Printf("11. logger.Printf(): Print log output in TRACE level\n")
	logger.Println("12. logger.Println(): Print log output in TRACE level")

	logger.Warning("13. logger.Waining(): WARNING log output in TRACE level\n")
	logger.Warningf("14. logger.Wainingf(): WARNING log output in TRACE level\n")
	logger.Warningln("15. logger.Wainingln(): WARNING log output in TRACE level")

	logger.Fatal("16. logger.Fatal(): FATAL log output in TRACE level\n")
	logger.Fatalf("17. logger.Fatalf(): FATAL log output in TRACE level\n")
	logger.Fatalln("18. logger.Fatalln(): FATAL log output in TRACE level")

	// Set level to DEBUG
	logger.SetMainLevel(LOG_LEVEL_DEBUG)
	logger.Debug("19. logger.Debug(): DEBUG log output in DEBUG level\n")
	logger.Debugf("20. logger.Debugf(): DEBUG log output in DEBUG level\n")
	logger.Debugln("21. logger.Debugln(): DEBUG log output in DEBUG level")

	logger.Print("21. logger.Print(): Print log output in DEBUG level\n")
	logger.Printf("22. logger.Printf(): Print log output in DEBUG level\n")
	logger.Println("23. logger.Println(): Print log output in DEBUG level")

	logger.Warning("24. logger.Waining(): WARNING log output in DEBUG level\n")
	logger.Warningf("25. logger.Wainingf(): WARNING log output in DEBUG level\n")
	logger.Warningln("26. logger.Wainingln(): WARNING log output in DEBUG level")

	logger.Fatal("27. logger.Fatal(): FATAL log output in DEBUG level\n")
	logger.Fatalf("28. logger.Fatalf(): FATAL log output in DEBUG level\n")
	logger.Fatalln("29. logger.Fatalln(): FATAL log output in DEBUG level")

	// Set level to FATAL
	logger.SetMainLevel(LOG_LEVEL_FATAL)
	logger.Debug("!!!!! logger.Debug(): no DEBUG log output in FATAL level\n")
	logger.Debugf("!!!!! logger.Debugf(): no DEBUG log output in FATAL level\n")
	logger.Debugln("!!!!! logger.Debugln(): no DEBUG log output in FATAL level")

	logger.Print("!!!!! logger.Print(): no Print log output in FATAL level\n")
	logger.Printf("!!!!! logger.Printf(): no Print log output in FATAL level\n")
	logger.Println("!!!!! logger.Println(): no Print log output in FATAL level")

	logger.Warning("!!!!! logger.Waining(): no WARNING log output in FATAL level\n")
	logger.Warningf("!!!!! logger.Wainingf(): no WARNING log output in FATAL level\n")
	logger.Warningln("!!!!! logger.Wainingln(): no WARNING log output in FATAL level")

	logger.Fatal("30. logger.Fatal(): FATAL log output in FATAL level\n")
	logger.Fatalf("31. logger.Fatalf(): FATAL log output in FATAL level\n")
	logger.Fatalln("32. logger.Fatalln(): FATAL log output in FATAL level")

	ma := logger.LoggerModule("ma")

	// Default level is FATAL
	logger.ModulePrint(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in default level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in default level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in default level")

	logger.ModulePrint(ma, LOG_LEVEL_TRACE, "!!!!! logger.ModulePrint(): no TRACE log output in default level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_TRACE, "!!!!! logger.ModulePrint(): no TRACE log output in default level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_TRACE, "!!!!! logger.ModulePrint(): no TRACE log output in default level")

	logger.ModulePrint(ma, LOG_LEVEL_WARNING, "!!!!! logger.ModulePrint(): no WARNING log output in default level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_WARNING, "!!!!! logger.ModulePrint(): no WARNING log output in default level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_WARNING, "!!!!! logger.ModulePrint(): no WARNING log output in default level")

	logger.ModulePrint(ma, LOG_LEVEL_FATAL, "33. logger.ModulePrint(): FATAL log output in default level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_FATAL, "34. logger.ModulePrint(): FATAL log output in default level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_FATAL, "35. logger.ModulePrint(): FATAL log output in default level")

	// Set level to WARNING
	logger.SetModuleLevel(ma, LOG_LEVEL_WARNING)
	logger.ModulePrint(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in WARNING level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in WARNING level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in WARNING level")

	logger.ModulePrint(ma, LOG_LEVEL_TRACE, "!!!!! logger.ModulePrint(): no TRACE log output in WARNING level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_TRACE, "!!!!! logger.ModulePrint(): no TRACE log output in WARNING level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_TRACE, "!!!!! logger.ModulePrint(): no TRACE log output in WARNING level")

	logger.ModulePrint(ma, LOG_LEVEL_WARNING, "36. logger.ModulePrint(): WARNING log output in WARNING level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_WARNING, "37. logger.ModulePrint(): WARNING log output in WARNING level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_WARNING, "38. logger.ModulePrint(): WARNING log output in WARNING level")

	logger.ModulePrint(ma, LOG_LEVEL_FATAL, "39. logger.ModulePrint(): FATAL log output in WARNING level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_FATAL, "40. logger.ModulePrint(): FATAL log output in WARNING level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_FATAL, "41. logger.ModulePrint(): FATAL log output in WARNING level")

	// Set level to TRACE
	logger.SetModuleLevel(ma, LOG_LEVEL_TRACE)
	logger.ModulePrint(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in TRACE level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in TRACE level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_DEBUG, "!!!!! logger.ModulePrint(): no DEBUG log output in TRACE level")

	logger.ModulePrint(ma, LOG_LEVEL_TRACE, "42. logger.ModulePrint(): TRACE log output in TRACE level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_TRACE, "43. logger.ModulePrint(): TRACE log output in TRACE level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_TRACE, "44. logger.ModulePrint(): TRACE log output in TRACE level")

	logger.ModulePrint(ma, LOG_LEVEL_WARNING, "45. logger.ModulePrint(): WARNING log output in TRACE level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_WARNING, "46. logger.ModulePrint(): WARNING log output in TRACE level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_WARNING, "47. logger.ModulePrint(): WARNING log output in TRACE level")

	logger.ModulePrint(ma, LOG_LEVEL_FATAL, "48. logger.ModulePrint(): FATAL log output in TRACE level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_FATAL, "49. logger.ModulePrint(): FATAL log output in TRACE level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_FATAL, "50. logger.ModulePrint(): FATAL log output in TRACE level")

	// Set level to DEBUG
	logger.SetModuleLevel(ma, LOG_LEVEL_DEBUG)
	logger.ModulePrint(ma, LOG_LEVEL_DEBUG, "51. logger.ModulePrint(): no DEBUG log output in DEBUG level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_DEBUG, "52. logger.ModulePrint(): no DEBUG log output in DEBUG level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_DEBUG, "53. logger.ModulePrint(): no DEBUG log output in DEBUG level")

	logger.ModulePrint(ma, LOG_LEVEL_TRACE, "54. logger.ModulePrint(): TRACE log output in DEBUG level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_TRACE, "55. logger.ModulePrint(): TRACE log output in DEBUG level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_TRACE, "56. logger.ModulePrint(): TRACE log output in DEBUG level")

	logger.ModulePrint(ma, LOG_LEVEL_WARNING, "57. logger.ModulePrint(): WARNING log output in DEBUG level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_WARNING, "58. logger.ModulePrint(): WARNING log output in DEBUG level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_WARNING, "59. logger.ModulePrint(): WARNING log output in DEBUG level")

	logger.ModulePrint(ma, LOG_LEVEL_FATAL, "60. logger.ModulePrint(): FATAL log output in DEBUG level\n")
	logger.ModulePrintf(ma, LOG_LEVEL_FATAL, "61. logger.ModulePrint(): FATAL log output in DEBUG level\n")
	logger.ModulePrintln(ma, LOG_LEVEL_FATAL, "62. logger.ModulePrint(): FATAL log output in DEBUG level")

	logger.Close()
	fmt.Println("+++++++++++++++++ TestLog ++++++++++++++++++++")
}
