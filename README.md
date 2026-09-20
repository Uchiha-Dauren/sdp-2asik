Logistics + GUI Demo
A small Go console application demonstrating two design patterns:
Factory Method — selects the delivery transport: Truck or Ship.
Abstract Factory — creates matching UI components for Windows or macOS.
The delivery method and UI platform are selected independently.
Invalid or missing arguments result in an error message and exit code 1. No default values are used.

Verification
The following cases were tested successfully:

ROAD WINDOWS
SEA WINDOWS
ROAD MACOS
SEA MACOS
Invalid delivery mode
Invalid platform
Missing argument# sdp-2asik
