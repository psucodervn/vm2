package pm2

import (
  "os/exec"
  "bytes"
  "github.com/pkg/errors"
  "encoding/json"
  "log"
)

type Monit struct {
  Memory int `json:"memory"`
  CPU    int `json:"cpu"`
}

type Env struct {
  ExitCode int    `json:"exit_code"`
  Status   string `json:"status"`
  ExecPath string `json:"pm_exec_path"`
  ExecMode string `json:"exec_mode"`
  Watch    bool   `json:"watch"`
}

type Process struct {
  Name string `json:"name"`
  PID  int    `json:"pid"`
  PMID int    `json:"pm_id"`
  Env  Env    `json:"pm2_env"`
}

var (
  ErrPM2PathNotSet = errors.New("pm2 executable path is not set")
)

func List() ([]Process, error) {
  cmd := exec.Command("pm2", "jlist")
  var stdout bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr

  if err := cmd.Run(); err != nil {
    log.Println(stderr.String())
    return nil, err
  }

  var procs []Process
  if err := json.Unmarshal(stdout.Bytes(), &procs); err != nil {
    return nil, err
  }

  return procs, nil
}
