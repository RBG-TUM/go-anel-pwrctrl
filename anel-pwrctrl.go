package go_anel_pwrctrl

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const indexOffset = 21 // index of first power outlet in status response

type PwrCtrl struct {
	addr  string
	auth  string
	state []string // see readme for documentation
}

//New creates an instance of PwrCtrl. Auth should be provided like "user:password"
func New(addr string, auth string) PwrCtrl {
	return PwrCtrl{addr: addr, auth: base64.StdEncoding.EncodeToString([]byte(auth))}
}

// TurnOn powers on the outlet with the index outletIndex
func (c *PwrCtrl) TurnOn(outletIndex int) error {
	return c.turn(outletIndex, true)
}

// TurnOff powers off the outlet with the index outletIndex
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

// IsOn returns whether the outlet with index outletIndex is on or not.
func (c *PwrCtrl) IsOn(outletIndex int) (bool, error) {
	err := c.updateStatus()
	if err != nil {
		return false, err
	}
	if len(c.state) < indexOffset+outletIndex {
		return false, fmt.Errorf("outlet index %d is out of range", outletIndex)
	}
	return c.state[20+outletIndex] == "1", nil
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
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	state := strings.Split(string(body), ";")
	// todo: perhaps sanity check response
	c.state = state
	return nil
}
