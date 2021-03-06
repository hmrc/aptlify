package action

import (
	"errors"
	"fmt"
	"strings"
)

func (a ActionStruct) Plan() error {

	var message string

	if a.isEmpty() {
		return errors.New("Uninitialised action")
	}

	switch {
	case a.ChangeType == MirrorUpdate:
		message = fmt.Sprintf("+mirror %s will be updated. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == MirrorCreate:
		message = fmt.Sprintf("+mirror %s will be created. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == MirrorRecreate:
		message = fmt.Sprintf("+/-mirror %s will be recreated. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == RepoCreate:
		message = fmt.Sprintf("+repo %s will be created. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == GpgAdd:
		message = fmt.Sprintf("+gpg key %s will be added. Reason: %s", a.ResourceName, strings.Join(a.changeReason, ","))
	case a.ChangeType == SnapshotUpdate:
		message = fmt.Sprintf("+snapshot %s will be updated at revision %05d. Reason: %s", a.ResourceName, a.SnapshotRevision, strings.Join(a.changeReason, ","))
	case a.ChangeType == Noop:
		message = fmt.Sprintf("resource unchanged: %s", a.ResourceName)
	default:
		message = fmt.Sprintf("no case matched")
	}

	fmt.Println(message)
	return nil

}
