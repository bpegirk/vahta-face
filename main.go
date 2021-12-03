package main

import (
	"bufio"
	"encoding/binary"
	_ "example.org/modules/config"
	_ "example.org/modules/database"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
	"net"
	"strings"
	"time"
)

const (
	EVENT_TURNIKET_ALLOW      = 21
	EVENT_CARD_DISABLED       = 52
	EVENT_DENY                = 54
	EVENT_DENY_DOUBLE_CONTROL = 58
	EVENT_CARDREADER          = 33
)

var conn net.Conn

func main() {
	initConfig()
	initSocket()
}

func initSocket() {
	// Подключаемся к сокету
	conString := fmt.Sprintf("%s:%v", configuration.Socket.Host, configuration.Socket.Port)
	conn, _ = net.Dial("tcp", conString)
	fmt.Println("Start listening...")
	fmt.Fprintf(conn, "\n")
	result := make([]byte, 128)
	for {
		_, err := bufio.NewReader(conn).Read(result)
		if err != nil {
			panic(err)
		}
		socketMessage(result)
	}
}

func socketMessage(result []byte) {
	stringBody := string(result)

	lastIndex := 0
	if strings.Contains(stringBody, "<START>") {
		lastIndex = strings.Index(stringBody, "<END>")
	} else {
	}

	var body []byte
	for i := 7; i < lastIndex; i++ {
		body = append(body, result[i])
	}
	fmt.Println("Raw data: ", body)
	parseBody(body)
}

func parseBody(body []byte) {
	messages := map[int]string{
		1:    "Доступ разрешен",
		3:    "Доступ разрешен, человек не прошел",
		5:    "Неизвестная карта",
		10:   "Доступ запрещен",
		11:   "Запрос на доступ",
		12:   "Предоставлен проход по кнопке",
		13:   "Проход по кнопке не состоялся",
		14:   "Проход по кнопке состоялся",
		16:   "Свободный проход",
		20:   "Дверь не закрыта",
		31:   "Перевод в режим 'Открыт' с пульта управления",
		33:   "Перевод в режим 'Дежурный'' с пульта управления",
		35:   "Разблокировка лифта по кнопке",
		36:   "Считыватели заблокированы",
		37:   "Считыватели разблокированы",
		51:   "Несанкционированное открытие двери",
		53:   "Дверь закрыта",
		100:  "Шлейф-тревога",
		101:  "Шлейф-на охране",
		102:  "Шлейф-готов",
		103:  "Шлейф-не готов",
		104:  "Шлейф разрушен (обрыв)",
		105:  "Шлейф разрушен (короткое замыкание)",
		111:  "Тампер корпуса-тревога",
		112:  "Тампер корпуса-на охране",
		150:  "Проблемы с питанием",
		161:  "Постоянное напряжение в норме",
		162:  "Переменое напряжение в норме",
		163:  "Постоянное напряжение ниже нормы",
		165:  "Переменое напряжение отсутствует",
		170:  "Включено питание",
		172:  "Контроллер переведён в режим программирования",
		173:  "Контроллер переведён в рабочий режим",
		180:  "Контроллер перевел часы на летнее время",
		181:  "Контроллер перевел часы на зимнее время",
		201:  "Связь отсутствует",
		202:  "Связь восстановлена",
		301:  "Группа снята с охраны",
		302:  "Группа поставлена на охрану",
		303:  "Отказ в управлении группой",
		304:  "Команда управления группой не выполнена",
		310:  "Группа поставлена на охрану по реакции",
		311:  "Группа снята с охраны по реакции",
		312:  "Группа поставлена на охрану оператором",
		313:  "Группа снята с охраны оператором",
		315:  "Группа снята с охраны под принуждением",
		316:  "Тревога в группе",
		434:  "Пожарный датчик - норма",
		435:  "Пожарный датчик - тревога",
		436:  "Отмена пожарной разблокировки",
		437:  "Пожарная разблокировка",
		500:  "Комментарий",
		501:  "Отключение пропуска",
		502:  "Включение пропуска",
		503:  "Отключение учетной записи",
		504:  "Включение учетной записи",
		505:  "Отключение доступа к Интернет",
		506:  "Включение доступа к Интернет",
		507:  "Удаление пропуска",
		508:  "Перерегистрация пропуска",
		509:  "Прикрепление пропуска",
		510:  "Изменение пропуска",
		511:  "Создание пользователя в FB",
		512:  "Контроль повторного входа включен",
		513:  "Контроль повторного входа выключен",
		514:  "Изменение пароля",
		515:  "Прикрепление документа",
		516:  "Открепление документа",
		517:  "Создание сетевого ресурса",
		518:  "Подключение сетевого ресурса",
		519:  "Отключение сетевого ресурса",
		520:  "Выдача интернет трафика",
		521:  "Изменение уровня доступа",
		522:  "Перенос между подразделениями",
		1001: "Команда управления не выполнена контроллером",
		1002: "Команда оператора",
		1003: "Вход оператора",
		1010: "Выдан ключ от помещения",
		1011: "Сдан ключ от помещения",
		1020: "Очистка памяти пропусков контроллера",
		1021: "Регистрация пропуска в памяти контроллера",
		1022: "Удаление пропуска из памяти контроллера",
		1023: "Изменение свойств пропуска в памяти контроллера",
		1050: "Сброс состояния по кнопке",
		1051: "Тамбур - тревога",
		1052: "Сброс состояния оператором",
		1053: "Аварийное открытие тамбура",
		1054: "Сброс состояния по таймауту",
		1055: "Перевод в ручной режим",
		2056: "counters",
		2057: "A ext busy",
		2058: "A int busy",
		2059: "B ext busy",
		2060: "B int busy",
		2061: "A int wait",
		2062: "B int wait",
	}

	command := int(binary.LittleEndian.Uint32(body))

	fmt.Println("Command:\t", command)

	allowEvents := []int{EVENT_TURNIKET_ALLOW, EVENT_CARDREADER, EVENT_CARD_DISABLED, EVENT_DENY_DOUBLE_CONTROL, EVENT_DENY}
	if indexOf(command, allowEvents) != -1 {
		pOffset := getPeopleOffset(command)
		hOffset := getHwOffset(command)
		mOffset := int(binary.LittleEndian.Uint32(part(body, 5, 4)))
		additionalMessage := ""
		if mOffset > 0 {
			additionalMessage = messages[mOffset]
		}
		currentTime := time.Now()

		//fmt.Println("eventText:\t", EventOffsets[EventCodes[eventId]].label)
		fmt.Println("messageId:\t", mOffset)
		fmt.Println("message:\t", additionalMessage)
		if hOffset > 0 {
			hardId := binary.LittleEndian.Uint32(part(body, hOffset, 4))
			fmt.Println("hard id:\t", hardId)
			if hardId > 0 {
				fmt.Println("hard name:\t", hardwareName(int(hardId)))

			}
		}
		if pOffset > 0 {
			userId := binary.LittleEndian.Uint32(part(body, pOffset, 4))
			fmt.Println("user id:\t", userId)
			if userId > 0 {
				fio, _ := peopleName(int(userId))
				fmt.Println("user name:\t", fio)

			}

		}
		fmt.Println("time:\t", currentTime.Format("15:04:05"))
		fmt.Println("data:\t", currentTime.Format("2006-01-02"))

	} else {
		fmt.Println("Not allowed command. Skip")
	}
	fmt.Println("-------------------------")
}

func part(body []byte, start int, len int) []byte {
	var res []byte
	for i := start; i < start+len; i++ {
		res = append(res, body[i])
	}
	return res
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
func getHwOffset(eventId int) int {
	switch eventId {
	case EVENT_CARDREADER, EVENT_CARD_DISABLED, EVENT_DENY_DOUBLE_CONTROL, EVENT_DENY:
		return 9
	case EVENT_TURNIKET_ALLOW:
		return 13
	}
	return 0
}
func getPeopleOffset(eventId int) int {
	switch eventId {
	case EVENT_CARDREADER, EVENT_CARD_DISABLED, EVENT_DENY_DOUBLE_CONTROL, EVENT_DENY:
		return 17
	case EVENT_TURNIKET_ALLOW:
		return 21
	}
	return 0
}
