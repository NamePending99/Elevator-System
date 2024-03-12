# Elevator-System

An elevator system with N elevators and M floors implemented in Go.

The network topology of the system is circular, where data flows from one elevator to the next.

## Program Notes

The program contains a elevator-state object (elevState) containing all neccesarry information. This object is only modified by reference when calling functions, meaning no function return is neccassary.

These functions do still return a pointer to the object, which functionally does nothing. The reason for this design is to clarify when an elevator-state object is modified.

## Installation

The elevator system only runs on Unix-like operating systems, and requires the elevator simulator from [Simulator-v2](https://github.com/TTK4145/Simulator-v2) to function correctly.

Install Simulator-v2 locally and follow the readme. If you are running macOS Sonoma, you will likely encounter a linker error when compiling the simulator with the given compile command. A quick fix can be retreived from [here](https://forum.dlang.org/thread/jwmpdecwyazcrxphttoy@forum.dlang.org).

Download the elevator-system with:

```bash
git clone git+https://github.com/Mathiasotnes/Deep-Learning.git
```

Cd to elevator-system and build the project with:

```bash
go build elevator
```

Create a Go folder in your /bin directory and add the build file there. If you are running on macOS, open a terminal and enter:

```bash
vim ~/.zshrc
```

Add the /bin/Go directroy to your PATH environment with:

```zsh
export PATH="${HOME}/bin/go:${PATH}"
```

Now you can start the elevator system on any terminal.

## Usage

Each elevator instance requires a Simulator instance. After completing the installation, open a terminal window and run `SimElevatorServer --port {your port here}` to start a simulator.

Open another terminal and enter `elevator -h`. This command lists alla required flags to start an elevator. The flags are:

- -id: node ID of the elevator (first elevator should be 0).
- -num: number of nodes (elevators) in the network.
- -bport: base-port for all elevators.
- -sport: server-port, the port the Simulator is running on.

### Example

Terminal 1:

```bash
SimElevatorServer --port 9090
```

Terminal 2:

```bash
elevator -id 0 -num 3 -bport 8080 -sport 9090
```

Repeat the commands above as many times as the -num flag. Remember to set -id, -sport and --port accordingly. This should yield a elevator system with 3 elevators and 4 floors. To use the Simulator, see the [readme](https://github.com/TTK4145/Simulator-v2).

## Repository activity

![Alt](https://repobeats.axiom.co/api/embed/359927114a85e3f95cbe1c54958fd3158a42f3e8.svg "Repobeats analytics image")
