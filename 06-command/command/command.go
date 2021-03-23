package command

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type NoCommand struct{}

func (nc NoCommand) Execute() {}
func (nc NoCommand) Undo()    {}

type Light struct {
}

func (l Light) on() {
	fmt.Println("Light is on")
}

func (l Light) off() {
	fmt.Println("Light is off")
}

type LightOnCommand struct {
	Light Light
}

func (loc LightOnCommand) Execute() {
	loc.Light.on()
}

func (loc LightOnCommand) Undo() {
	loc.Light.off()
}

type LightOffCommand struct {
	Light Light
}

func (lof LightOffCommand) Execute() {
	lof.Light.off()
}

func (lof LightOffCommand) Undo() {
	lof.Light.on()
}

type SimpleRemoteControl struct {
	slot Command
}

func (src *SimpleRemoteControl) SetCommand(command Command) {
	src.slot = command
}

func (src SimpleRemoteControl) ButtonWasPressed() {
	src.slot.Execute()
}

type RemoteControl struct {
	onCommands  [7]Command
	offCommands [7]Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	var onCommands [7]Command
	var offCommands [7]Command
	noCommand := new(NoCommand)
	for i := 0; i < 7; i++ {
		onCommands[i] = noCommand
		offCommands[i] = noCommand
	}
	undoCommand := noCommand
	rc := &RemoteControl{onCommands: onCommands, offCommands: offCommands, undoCommand: undoCommand}
	return rc
}

func (rc *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	rc.onCommands[slot] = onCommand
	rc.offCommands[slot] = offCommand
}

func (rc *RemoteControl) OnButtonWasPushed(slot int) {
	rc.onCommands[slot].Execute()
	rc.undoCommand = rc.onCommands[slot]
}

func (rc *RemoteControl) OffButtonWasPushed(slot int) {
	rc.offCommands[slot].Execute()
	rc.undoCommand = rc.offCommands[slot]
}

func (rc *RemoteControl) UndoButtonWasPushed() {
	rc.undoCommand.Undo()
}

func (rc *RemoteControl) String() string {
	var str string
	for i := 0; i < 7; i++ {
		str += fmt.Sprintf("[slot %d ]  %#v %#v\n", i, rc.onCommands[i], rc.offCommands[i])
	}
	return str
}
