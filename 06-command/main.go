package main

import (
	"fmt"

	"github.com/design-pattern/06-command/command"
)

func main() {
	remote := command.RemoteControl{}

	livingRoomLight := command.Light{}
	kitchenRoomLight := command.Light{}

	livingRoomLightOn := command.LightOnCommand{Light: livingRoomLight}
	livingRoomLightOff := command.LightOffCommand{Light: livingRoomLight}

	kitchenRoomLightOn := command.LightOnCommand{Light: kitchenRoomLight}
	kitchenRoomLightOff := command.LightOffCommand{Light: kitchenRoomLight}

	remote.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remote.SetCommand(1, kitchenRoomLightOn, kitchenRoomLightOff)

	fmt.Println(remote.String())

	remote.OnButtonWasPushed(0)
	remote.OffButtonWasPushed(0)
	remote.OnButtonWasPushed(1)
	remote.OffButtonWasPushed(1)

	// [slot 0 ]  command.LightOnCommand{Light:command.Light{}} command.LightOffCommand{Light:command.Light{}}
	// [slot 1 ]  command.LightOnCommand{Light:command.Light{}} command.LightOffCommand{Light:command.Light{}}
	// [slot 2 ]  <nil> <nil>
	// [slot 3 ]  <nil> <nil>
	// [slot 4 ]  <nil> <nil>
	// [slot 5 ]  <nil> <nil>
	// [slot 6 ]  <nil> <nil>

	// Light is on
	// Light is off
	// Light is on
	// Light is off

	remote.OnButtonWasPushed(0)
	remote.OffButtonWasPushed(0)
	fmt.Println(remote.String())

	// 실행 취소 전등을 다시 킴
	remote.UndoButtonWasPushed()
	remote.OffButtonWasPushed(0)
	remote.OnButtonWasPushed(0)
	fmt.Println(remote.String())

	// 실행 취소 전등을 다시 끔
	remote.UndoButtonWasPushed()

	// Light is on
	// Light is off
	// [slot 0 ]  command.LightOnCommand{Light:command.Light{}} command.LightOffCommand{Light:command.Light{}}
	// [slot 1 ]  command.LightOnCommand{Light:command.Light{}} command.LightOffCommand{Light:command.Light{}}
	// [slot 2 ]  <nil> <nil>
	// [slot 3 ]  <nil> <nil>
	// [slot 4 ]  <nil> <nil>
	// [slot 5 ]  <nil> <nil>
	// [slot 6 ]  <nil> <nil>

	// Light is on
	// Light is off
	// Light is on
	// [slot 0 ]  command.LightOnCommand{Light:command.Light{}} command.LightOffCommand{Light:command.Light{}}
	// [slot 1 ]  command.LightOnCommand{Light:command.Light{}} command.LightOffCommand{Light:command.Light{}}
	// [slot 2 ]  <nil> <nil>
	// [slot 3 ]  <nil> <nil>
	// [slot 4 ]  <nil> <nil>
	// [slot 5 ]  <nil> <nil>
	// [slot 6 ]  <nil> <nil>

	// Light is off
}
