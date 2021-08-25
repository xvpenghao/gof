package command

import "testing"

func TestCmd(t *testing.T) {
	var tv Device = &TV{} // 接受者

	// 命令
	var oncmd Cmd = &OnCmd{Dev: tv}
	var offcmd Cmd = &OffCmd{Dev: tv}

	// 调用者
	btn := &Button{Cmd: oncmd}
	btn.Press()

	btn = &Button{Cmd: offcmd}
	btn.Press()

}
