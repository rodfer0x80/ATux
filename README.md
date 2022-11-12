# ATux
-----------

## Idea
-----------
We want to accept any arcade and retro game and controller and run it max
performance with modern graphical options.
Thinking set_supported_consoles<=PS1

## Tools
-----------
-- Software
- C, Rust, Haskell
- Assembly
- FPGA verilog RISC-V
-- Hardware
- RISC-V FPGA
- Screen display (development)
- ESP-32 Wifi chip
- Mobile phone running controller ap (development)
- Controller app (to be developed)
- Power supply and charger
- Motherboard
- USB-C bus
- RAM
- SSD Disk Local
- SSD Disk Storage

## Flow:
-----------
. [USER] Controller interface ->
. [USER] Physical Medium ->
. Parser module ->
. Transpiller module ->
. Runner module ->
. [USER] Game Display Module
