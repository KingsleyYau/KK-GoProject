package pushserver

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"time"
)

func LogFile(logString string) {
	os.Mkdir("log", 0766)
	DailyLogToFile("log/push_server", logString)
}
func DailyLogToFile(filePreName string, logString string) {
	newFileName := filePreName + "_" + time.Now().Format("2006-01-02") + ".log"
	logFile,err := os.OpenFile(newFileName, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0766);
    if err != nil {
        log.Fatal("%s\r\n", err.Error())
    }
    defer logFile.Close();
    logger := log.New(logFile, "\r\n", log.LstdFlags);
    logger.Println(logString);
}
func Md5tohex(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	hexString := hex.EncodeToString(h.Sum(nil))
	return hexString
}