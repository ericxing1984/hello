package logger

import (
	"fmt"

	log "github.com/cihub/seelog"
)

func TestLog() {
	fmt.Println("Test log")
	defer log.Flush()

	//替换记录器
	log.ReplaceLogger(getCfgLogger())

	log.Info("Hello world from Seelog!")
	log.Info("Hello world from Seelog!")
	//dateTimeFormat()
}

func getXmlLogger() log.LoggerInterface {
	//加载配置文件
	logger, err := log.LoggerFromConfigAsFile("./logger/config.xml")

	if err != nil {

		fmt.Println("parse config.xml error")
	}
	return logger
}

func getCfgLogger() log.LoggerInterface {
	logCfg := LogCfg{"info", "main", "date", "test.log", "02.01.2006", "7"}
	logger, err := RollingFile(logCfg)
	if err != nil {
		fmt.Println(err)
	}
	return logger
}

type LogCfg struct {
	MinLevel            string `json:"min_level"`
	Fmt_id              string `json:"fmt_id"`
	Roll_type           string `json:"roll_type"`
	File_name           string `json:"file_name"`
	Roll_type_param     string `json:"roll_type_param"`
	Roll_type_max_rolls string `json:"roll_type_max_rolls"`
}

type LocalLogInterface log.LoggerInterface

func RollingFile(cfg LogCfg) (LocalLogInterface, error) {

	logConfig_Temp := `
<seelog type="asynctimer" asyncinterval="5000000" minlevel="%s">  
	<outputs formatid="common">  
 
	<buffered formatid="common" size="%d" flushperiod="1000">  
 
	<rollingfile type="date" filename="%s%s" %s="%s" fullname="true" maxrolls="%d"/>  
	</buffered>  
 
	</outputs>  
 
	 <formats>  
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />  
	 </formats>  
 </seelog>  `

	//cfg_roll_type_map := map[string]string{"date": "datepattern", "size": "maxsize"}

	//roll_type_param_key := cfg_roll_type_map[cfg.Roll_type]

	/* 	logConfig := fmt.Sprintf(logConfig_Temp,
	cfg.Levels,
	cfg.Fmt_id,
	cfg.Roll_type,
	cfg.File_name,
	roll_type_param_key,
	cfg.Roll_type_param,
	cfg.Roll_type_max_rolls) */

	logger, err := log.LoggerFromConfigAsBytes([]byte(logConfig_Temp))

	if err != nil {
		fmt.Println(err)
	}

	return logger, err
}

func dateTimeFormat() {
	fmt.Println("Date time format")

	testConfig := `
	<seelog type="asynctimer" asyncinterval="5000000" minlevel="info">  
	<outputs formatid="common">  
 
	<buffered formatid="common" size="10000" flushperiod="1000">  
 
	<rollingfile type="date" filename="./log/test.log" datepattern="02.01.2006" fullname="true" maxrolls="30"/>  
	</buffered>  
 
	</outputs>  
 
	 <formats>  
		 <format id="common" format="%Date %Time [%LEV] [%File:%Line] [%Func] %Msg%n" />  
	 </formats>  
 </seelog>  `

	logger, err := log.LoggerFromConfigAsBytes([]byte(testConfig))

	if err != nil {
		fmt.Println(err)
	}

	loggerErr := log.ReplaceLogger(logger)

	if loggerErr != nil {
		fmt.Println(loggerErr)
	}

	log.Info("Test message!")
}
