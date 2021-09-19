package go_anel_pwrctrl

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type PwrCtrl struct {
	addr  string
	auth  string
	state []string // see readme for documentation
}

//New creates an instance of PwrCtrl. Auth should be provided in the like "user:password"
func New(addr string, auth string) PwrCtrl {
	return PwrCtrl{addr: addr, auth: base64.StdEncoding.EncodeToString([]byte(auth))}
}

func (c *PwrCtrl) TurnOn(outletIndex int) error {
	return c.turn(outletIndex, true)
}

func (c *PwrCtrl) TurnOff(outletIndex int) error {
	return c.turn(outletIndex, false)
}

func (c *PwrCtrl) turn(outletIndex int, targetOn bool) error {
	on, err := c.IsOn(outletIndex)
	if err != nil {
		return err
	}
	if on != targetOn {
		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/ctrl.htm", c.addr), strings.NewReader(fmt.Sprintf("F%d=S", outletIndex)))
		if err != nil {
			return err
		}
		request.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.auth))
		_, err = http.DefaultClient.Do(request)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *PwrCtrl) IsOn(index int) (bool, error) {
	err := c.updateStatus()
	if err != nil {
		return false, err
	}
	return c.state[20+index] == "1", nil
}

func (c *PwrCtrl) updateStatus() error {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/strg.cfg", c.addr), nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.auth))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	state := strings.Split(string(body), ";")
	if len(state) < 58 || state[58] != "end" {
		return errors.New("unexpected response from device")
	}
	c.state = state
	return nil
}
