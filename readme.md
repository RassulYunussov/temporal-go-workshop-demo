# Temporal workshop

The following example shows how to use Temporal Framework to implement simple workflow of money transfer business case

# Prerequisites

You need to install brew formula: https://formulae.brew.sh/formula/temporal
You need to start local Temporal Server: 

`temporal server start-dev`

# Project scope

Application initiates worker & wofkflows generator to simulate money transfer.
All operations are orchestrated by Temporal Framework. It is responsible to finsih the process attempting to repeat operations from the state it was interrupted.