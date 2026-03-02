

# Critterarium Proposal:

**Team Members:**

* Jack Werremeyer
* Nolan


## 1. Project Summary

Critterarium is a ecologoical simulation full of virtual critters that allows for users add food and critters to the enviroment. Each critter can move, eat, react to neighbors, and reproduces inside a shared simulated world.

Hopefully, by the end of the project, the system will resemble an ecosystem, where interactions can result in fun and interesting behaviors


## 2. Typical Use Cases

1. Running the ecosystem

The backend server simulates the aquarium world, running thousands of critter programs concurrently. Critters interact with each other and the environment in real time.


2. GUI Monitoring and Control

Users connect to the simulation using a GUI client to visualize critters, pause or resume the simulation, inspect individual programs, and adjust parameters like food density or reproduction rate. (maybe add predators too)



## 3. Intended Design and Components

The project will be organized into multiple components:

### Core Modules

1. critter/internal

   * Contains the core critter functionality

2. critter/controller

   * Controls critter actions and keeps the current state of a critter


### Simulation/System Modules

3. simulation/world

   * Aquarium grid or continuous space that handles physics, food spawning, collisions, and interactions

4. simulation/server

   * Multi-threaded Go server that runs the simulation loop and critter controllers

5. client/gui

    * GUI for visualizing the aquarium and controlling the simulation


## 4. Testing Strategy

Testing will largely focus on ensuring the critters function as intended and can properly interact with their enviroment. Also, we'll be testing the GUI for visual bug



## 5. Minimal Viable Product 

The MVP will include:

* Working Critters who can move, reproduce, and eat
* An enviroment where multiple critters can move and interact in
* Basic user input
* A command-line/minimal visualization of the simulation
* Necessary Unit tests for the above



## 6. Stretch Goals

* A GUI visualization (maybe aquarium themed?)
* A fault injector for critter mutations
* Distributed simulation across multiple servers
* More advanced Critter features


## 7. Expected Functionality by Checkpoint

By the checkpoint, the project is expected to include:

* Working critter functionality
* basic structure of the enviroment
* Unit tests
* Basic scaffolding for the simulation server
