package main

import (
	"fmt"
	"bytes"
	"log"
	"encoding/binary"
)

type projectDef uint8 
type classDef uint8
type cmdDef uint16

type command struct {
	project projectDef
	class   classDef
	cmd     cmdDef
}

// All ARDrone3-only commands
const projectardrone3 projectDef = 1
// All commands related to piloting the drone
const classPiloting classDef = 0
// title : Take off, 
// desc : Ask the drone to take off.\n On the fixed wings (such as Disco): not used except to cancel a land., 
// support : 0901;090c;090e, 
// result : On the quadcopters: the drone takes off if its [FlyingState](#1-4-1) was landed.\n On the fixed wings, the landing process is aborted if the [FlyingState](#1-4-1) was landing.\n Then, event [FlyingState](#1-4-1) is triggered., 
const cmdTakeOff cmdDef = 1

type ardrone3PilotingTakeOff command

type ardrone3PilotingTakeOffArguments struct {
}

func (a ardrone3PilotingTakeOff) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingTakeOffArguments{}
// No arguments to decode here !!

return arg
}

var pilotingTakeOff = ardrone3PilotingTakeOff {
project: projectardrone3,
class: classPiloting,
cmd: cmdTakeOff,
}

// title : Move the drone, 
// desc : Move the drone.\n The libARController is sending the command each 50ms.\n\n **Please note that you should call setPilotingPCMD and not sendPilotingPCMD because the libARController is handling the periodicity and the buffer on which it is sent.**, 
// support : 0901;090c;090e, 
// result : The drone moves! Yaaaaay!\n Event [SpeedChanged](#1-4-5), [AttitudeChanged](#1-4-6) and [PositionChanged](#1-4-4) (only if gps of the drone has fixed) are triggered., 
const cmdPCMD cmdDef = 2

type ardrone3PilotingPCMD command

type ardrone3PilotingPCMDArguments struct {
flag uint8
roll int8
pitch int8
yaw int8
gaz int8
timestampAndSeqNum uint32
}

func (a ardrone3PilotingPCMD) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingPCMDArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.flag)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.roll)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.pitch)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.yaw)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.gaz)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.timestampAndSeqNum)
offset += 4

return arg
}

var pilotingPCMD = ardrone3PilotingPCMD {
project: projectardrone3,
class: classPiloting,
cmd: cmdPCMD,
}

// title : Land, 
// desc : Land.\n Please note that on copters, if you put some positive gaz (in the [PilotingCommand](#1-0-2)) during the landing, it will cancel it., 
// support : 0901;090c;090e, 
// result : On the copters, the drone lands if its [FlyingState](#1-4-1) was taking off, hovering or flying.\n On the fixed wings, the drone lands if its [FlyingState](#1-4-1) was hovering or flying.\n Then, event [FlyingState](#1-4-1) is triggered., 
const cmdLanding cmdDef = 3

type ardrone3PilotingLanding command

type ardrone3PilotingLandingArguments struct {
}

func (a ardrone3PilotingLanding) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingLandingArguments{}
// No arguments to decode here !!

return arg
}

var pilotingLanding = ardrone3PilotingLanding {
project: projectardrone3,
class: classPiloting,
cmd: cmdLanding,
}

// title : Cut out the motors, 
// desc : Cut out the motors.\n This cuts immediatly the motors. The drone will fall.\n This command is sent on a dedicated high priority buffer which will infinitely retry to send it if the command is not delivered., 
// support : 0901;090c;090e, 
// result : The drone immediatly cuts off its motors.\n Then, event [FlyingState](#1-4-1) is triggered., 
const cmdEmergency cmdDef = 4

type ardrone3PilotingEmergency command

type ardrone3PilotingEmergencyArguments struct {
}

func (a ardrone3PilotingEmergency) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingEmergencyArguments{}
// No arguments to decode here !!

return arg
}

var pilotingEmergency = ardrone3PilotingEmergency {
project: projectardrone3,
class: classPiloting,
cmd: cmdEmergency,
}

// title : Return home, 
// desc : Return home.\n Ask the drone to fly to its [HomePosition](#1-24-0).\n The availability of the return home can be get from [ReturnHomeState](#1-4-3).\n Please note that the drone will wait to be hovering to start its return home. This means that it will wait to have a [flag](#1-0-2) set at 0., 
// support : 0901;090c;090e, 
// result : The drone will fly back to its home position.\n Then, event [ReturnHomeState](#1-4-3) is triggered.\n You can get a state pending if the drone is not ready to start its return home process but will do it as soon as it is possible., 
const cmdNavigateHome cmdDef = 5

type ardrone3PilotingNavigateHome command

type ardrone3PilotingNavigateHomeArguments struct {
start uint8
}

func (a ardrone3PilotingNavigateHome) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingNavigateHomeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.start)
offset++ 

return arg
}

var pilotingNavigateHome = ardrone3PilotingNavigateHome {
project: projectardrone3,
class: classPiloting,
cmd: cmdNavigateHome,
}

// title : Auto take off mode, 
// desc : Auto take off mode., 
const cmdAutoTakeOffMode cmdDef = 6

type ardrone3PilotingAutoTakeOffMode command

type ardrone3PilotingAutoTakeOffModeArguments struct {
state uint8
}

func (a ardrone3PilotingAutoTakeOffMode) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingAutoTakeOffModeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.state)
offset++ 

return arg
}

var pilotingAutoTakeOffMode = ardrone3PilotingAutoTakeOffMode {
project: projectardrone3,
class: classPiloting,
cmd: cmdAutoTakeOffMode,
}

// title : Move the drone to a relative position, 
// desc : Move the drone to a relative position and rotate heading by a given angle.\n Moves are relative to the current drone orientation, (drone's reference).\n Also note that the given rotation will not modify the move (i.e. moves are always rectilinear)., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The drone will move of the given offsets.\n Then, event [RelativeMoveEnded](#1-34-0) is triggered.\n If you send a second relative move command, the drone will trigger a [RelativeMoveEnded](#1-34-0) with the offsets it managed to do before this new command and the value of error set to interrupted., 
const cmdmoveBy cmdDef = 7

type ardrone3PilotingmoveBy command

type ardrone3PilotingmoveByArguments struct {
dX float32
dY float32
dZ float32
dPsi float32
}

func (a ardrone3PilotingmoveBy) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingmoveByArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dX)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dY)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dZ)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dPsi)
offset += 4

return arg
}

var pilotingmoveBy = ardrone3PilotingmoveBy {
project: projectardrone3,
class: classPiloting,
cmd: cmdmoveBy,
}

// title : Prepare the drone to take off, 
// desc : Prepare the drone to take off.\n On copters: initiates the thrown takeoff. Note that the drone will do the thrown take off even if it is steady.\n On fixed wings: initiates the take off process on the fixed wings.\n\n Setting the state to 0 will cancel the preparation. You can cancel it before that the drone takes off., 
// support : 090e;090c:4.3.0, 
// result : The drone will arm its motors if not already armed.\n Then, event [FlyingState](#1-4-1) is triggered with state set at motor ramping.\n Then, event [FlyingState](#1-4-1) is triggered with state set at userTakeOff.\n Then user can throw the drone to make it take off., 
const cmdUserTakeOff cmdDef = 8

type ardrone3PilotingUserTakeOff command

type ardrone3PilotingUserTakeOffArguments struct {
state uint8
}

func (a ardrone3PilotingUserTakeOff) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingUserTakeOffArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.state)
offset++ 

return arg
}

var pilotingUserTakeOff = ardrone3PilotingUserTakeOff {
project: projectardrone3,
class: classPiloting,
cmd: cmdUserTakeOff,
}

// title : Circle, 
// desc : Make the fixed wing circle.\n The circle will use the [CirclingAltitude](#1-6-14) and the [CirclingRadius](#1-6-13), 
// support : 090e, 
// result : The fixed wing will circle in the given direction.\n Then, event [FlyingState](#1-4-1) is triggered with state set at hovering., 
const cmdCircle cmdDef = 9

type ardrone3PilotingCircle command

type ardrone3PilotingCircleArguments struct {
direction uint32
}

func (a ardrone3PilotingCircle) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingCircleArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.direction)
offset += 4

return arg
}

var pilotingCircle = ardrone3PilotingCircle {
project: projectardrone3,
class: classPiloting,
cmd: cmdCircle,
}

// title : Move to a location, 
// desc : Move the drone to a specified location.\n If a new command moveTo is sent, the drone will immediatly run it (no cancel will be issued).\n If a [CancelMoveTo](#1-0-11) command is sent, the moveTo is stopped.\n During the moveTo, all pitch, roll and gaz values of the piloting command will be ignored by the drone.\n However, the yaw value can be used., 
// support : 090c:4.3.0, 
// result : Event [MovingTo](#1-4-12) is triggered with state running. Then, the drone will move to the given location.\n Then, event [MoveToChanged](#1-4-12) is triggered with state succeed., 
const cmdmoveTo cmdDef = 10

type ardrone3PilotingmoveTo command

type ardrone3PilotingmoveToArguments struct {
latitude float64
longitude float64
altitude float64
orientationmode uint32
heading float32
}

func (a ardrone3PilotingmoveTo) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingmoveToArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.orientationmode)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.heading)
offset += 4

return arg
}

var pilotingmoveTo = ardrone3PilotingmoveTo {
project: projectardrone3,
class: classPiloting,
cmd: cmdmoveTo,
}

// title : Cancel the moveTo, 
// desc : Cancel the current moveTo.\n If there is no current moveTo, this command has no effect., 
// support : 090c:4.3.0, 
// result : Event [MoveToChanged](#1-4-12) is triggered with state canceled., 
const cmdCancelMoveTo cmdDef = 11

type ardrone3PilotingCancelMoveTo command

type ardrone3PilotingCancelMoveToArguments struct {
}

func (a ardrone3PilotingCancelMoveTo) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingCancelMoveToArguments{}
// No arguments to decode here !!

return arg
}

var pilotingCancelMoveTo = ardrone3PilotingCancelMoveTo {
project: projectardrone3,
class: classPiloting,
cmd: cmdCancelMoveTo,
}

// title : Start a piloted POI, 
// desc : Start a piloted Point Of Interest.\n During a piloted POI, the drone will always look at the given POI but can be piloted normally. However, yaw value is ignored. Camera tilt and pan command is also ignored.\n Ignored if [PilotedPOI](#1-4-14) state is UNAVAILABLE., 
// support : 090c:4.3.0, 
// result : If the drone is hovering, event [PilotedPOI](#1-4-14) is triggered with state RUNNING. If the drone is not hovering, event [PilotedPOI](#1-4-14) is triggered with state PENDING, waiting to hover. When the drone hovers, the state will change to RUNNING. If the drone does not hover for a given time, piloted POI is canceled by the drone and state will change to AVAILABLE. Then, the drone will look at the given location., 
const cmdStartPilotedPOI cmdDef = 12

type ardrone3PilotingStartPilotedPOI command

type ardrone3PilotingStartPilotedPOIArguments struct {
latitude float64
longitude float64
altitude float64
}

func (a ardrone3PilotingStartPilotedPOI) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStartPilotedPOIArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8

return arg
}

var pilotingStartPilotedPOI = ardrone3PilotingStartPilotedPOI {
project: projectardrone3,
class: classPiloting,
cmd: cmdStartPilotedPOI,
}

// title : Stop the piloted POI, 
// desc : Stop the piloted Point Of Interest.\n If [PilotedPOI](#1-4-14) state is RUNNING or PENDING, stop it., 
// support : 090c:4.3.0, 
// result : Event [PilotedPOI](#1-4-14) is triggered with state AVAILABLE., 
const cmdStopPilotedPOI cmdDef = 13

type ardrone3PilotingStopPilotedPOI command

type ardrone3PilotingStopPilotedPOIArguments struct {
}

func (a ardrone3PilotingStopPilotedPOI) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStopPilotedPOIArguments{}
// No arguments to decode here !!

return arg
}

var pilotingStopPilotedPOI = ardrone3PilotingStopPilotedPOI {
project: projectardrone3,
class: classPiloting,
cmd: cmdStopPilotedPOI,
}

// title : Cancel the relative move, 
// desc : Cancel the current relative move.\n If there is no current relative move, this command has no effect., 
// result : Event [RelativeMoveChanged](#1-4-16) is triggered with state canceled., 
const cmdCancelMoveBy cmdDef = 14

type ardrone3PilotingCancelMoveBy command

type ardrone3PilotingCancelMoveByArguments struct {
}

func (a ardrone3PilotingCancelMoveBy) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingCancelMoveByArguments{}
// No arguments to decode here !!

return arg
}

var pilotingCancelMoveBy = ardrone3PilotingCancelMoveBy {
project: projectardrone3,
class: classPiloting,
cmd: cmdCancelMoveBy,
}

// Animation commands
const classAnimations classDef = 5
// title : Make a flip, 
// desc : Make a flip., 
// support : 0901;090c, 
// result : The drone will make a flip if it has enough battery., 
const cmdFlip cmdDef = 0

type ardrone3AnimationsFlip command

type ardrone3AnimationsFlipArguments struct {
direction uint32
}

func (a ardrone3AnimationsFlip) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3AnimationsFlipArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.direction)
offset += 4

return arg
}

var animationsFlip = ardrone3AnimationsFlip {
project: projectardrone3,
class: classAnimations,
cmd: cmdFlip,
}

// Ask the drone to move camera
const classCamera classDef = 1
// title : Move the camera, 
// desc : Move the camera.\n You can get min and max values for tilt and pan using [CameraInfo](#0-15-0)., 
// support : 0901;090c;090e, 
// result : The drone moves its camera.\n Then, event [CameraOrientation](#1-25-0) is triggered., 
const cmdOrientation cmdDef = 0

type ardrone3CameraOrientation command

type ardrone3CameraOrientationArguments struct {
tilt int8
pan int8
}

func (a ardrone3CameraOrientation) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraOrientationArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.tilt)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.pan)
offset++ 

return arg
}

var cameraOrientation = ardrone3CameraOrientation {
project: projectardrone3,
class: classCamera,
cmd: cmdOrientation,
}

// title : Move the camera, 
// desc : Move the camera.\n You can get min and max values for tilt and pan using [CameraInfo](#0-15-0)., 
// support : 0901;090c;090e, 
// result : The drone moves its camera.\n Then, event [CameraOrientationV2](#1-25-2) is triggered., 
const cmdOrientationV2 cmdDef = 1

type ardrone3CameraOrientationV2 command

type ardrone3CameraOrientationV2Arguments struct {
tilt float32
pan float32
}

func (a ardrone3CameraOrientationV2) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraOrientationV2Arguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.tilt)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.pan)
offset += 4

return arg
}

var cameraOrientationV2 = ardrone3CameraOrientationV2 {
project: projectardrone3,
class: classCamera,
cmd: cmdOrientationV2,
}

// title : Move the camera using velocity, 
// desc : Move the camera given velocity consign.\n You can get min and max values for tilt and pan using [CameraVelocityRange](#1-25-4)., 
// support : 0901;090c;090e, 
// result : The drone moves its camera.\n Then, event [CameraOrientationV2](#1-25-2) is triggered., 
const cmdVelocity cmdDef = 2

type ardrone3CameraVelocity command

type ardrone3CameraVelocityArguments struct {
tilt float32
pan float32
}

func (a ardrone3CameraVelocity) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraVelocityArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.tilt)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.pan)
offset += 4

return arg
}

var cameraVelocity = ardrone3CameraVelocity {
project: projectardrone3,
class: classCamera,
cmd: cmdVelocity,
}

// Media recording management
const classMediaRecord classDef = 7
// title : Take a picture, 
// desc : Take a picture., 
const cmdPicture cmdDef = 0

type ardrone3MediaRecordPicture command

type ardrone3MediaRecordPictureArguments struct {
massstorageid uint8
}

func (a ardrone3MediaRecordPicture) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordPictureArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.massstorageid)
offset++ 

return arg
}

var mediaRecordPicture = ardrone3MediaRecordPicture {
project: projectardrone3,
class: classMediaRecord,
cmd: cmdPicture,
}

// title : Record a video, 
// desc : Record a video., 
const cmdVideo cmdDef = 1

type ardrone3MediaRecordVideo command

type ardrone3MediaRecordVideoArguments struct {
record uint32
massstorageid uint8
}

func (a ardrone3MediaRecordVideo) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordVideoArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.record)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.massstorageid)
offset++ 

return arg
}

var mediaRecordVideo = ardrone3MediaRecordVideo {
project: projectardrone3,
class: classMediaRecord,
cmd: cmdVideo,
}

// title : Take a picture, 
// desc : Take a picture.\n The type of picture taken is related to the picture setting.\n You can set the picture format by sending the command [SetPictureFormat](#1-19-0). You can also get the current picture format with [PictureFormat](#1-20-0).\n Please note that the time required to take the picture is highly related to this format.\n\n You can check if the picture taking is available with [PictureState](#1-8-2).\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after that the picture has been taken)., 
// support : 0901:2.0.1;090c;090e, 
// result : Event [PictureState](#1-8-2) will be triggered with a state busy.\n The drone will take a picture.\n Then, when picture has been taken, notification [PictureEvent](#1-3-0) is triggered.\n And normally [PictureState](#1-8-2) will be triggered with a state ready., 
const cmdPictureV2 cmdDef = 2

type ardrone3MediaRecordPictureV2 command

type ardrone3MediaRecordPictureV2Arguments struct {
}

func (a ardrone3MediaRecordPictureV2) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordPictureV2Arguments{}
// No arguments to decode here !!

return arg
}

var mediaRecordPictureV2 = ardrone3MediaRecordPictureV2 {
project: projectardrone3,
class: classMediaRecord,
cmd: cmdPictureV2,
}

// title : Record a video, 
// desc : Record a video (or start timelapse).\n You can check if the video recording is available with [VideoState](#1-8-3).\n This command can start a video (obvious huh?), but also a timelapse if the timelapse mode is set. You can check if the timelapse mode is set with the event [TimelapseMode](#1-20-4).\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after the picture has been taken)., 
// support : 0901:2.0.1;090c;090e, 
// result : The drone will begin or stop to record the video (or timelapse).\n Then, event [VideoState](#1-8-3) will be triggered. Also, notification [VideoEvent](#1-3-1) is triggered., 
const cmdVideoV2 cmdDef = 3

type ardrone3MediaRecordVideoV2 command

type ardrone3MediaRecordVideoV2Arguments struct {
record uint32
}

func (a ardrone3MediaRecordVideoV2) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordVideoV2Arguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.record)
offset += 4

return arg
}

var mediaRecordVideoV2 = ardrone3MediaRecordVideoV2 {
project: projectardrone3,
class: classMediaRecord,
cmd: cmdVideoV2,
}

// State of media recording
const classMediaRecordState classDef = 8
// title : Picture state, 
// desc : Picture state., 
const cmdPictureStateChanged cmdDef = 0

type ardrone3MediaRecordStatePictureStateChanged command

type ardrone3MediaRecordStatePictureStateChangedArguments struct {
state uint8
massstorageid uint8
}

func (a ardrone3MediaRecordStatePictureStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordStatePictureStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.state)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.massstorageid)
offset++ 

return arg
}

var mediaRecordStatePictureStateChanged = ardrone3MediaRecordStatePictureStateChanged {
project: projectardrone3,
class: classMediaRecordState,
cmd: cmdPictureStateChanged,
}

// title : Video record state, 
// desc : Picture record state., 
const cmdVideoStateChanged cmdDef = 1

type ardrone3MediaRecordStateVideoStateChanged command

type ardrone3MediaRecordStateVideoStateChangedArguments struct {
state uint32
massstorageid uint8
}

func (a ardrone3MediaRecordStateVideoStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordStateVideoStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.massstorageid)
offset++ 

return arg
}

var mediaRecordStateVideoStateChanged = ardrone3MediaRecordStateVideoStateChanged {
project: projectardrone3,
class: classMediaRecordState,
cmd: cmdVideoStateChanged,
}

// title : Picture state, 
// desc : Picture state., 
// support : 0901:2.0.1;090c;090e, 
// triggered : by [TakePicture](#1-7-2) or by a change in the picture state, 
const cmdPictureStateChangedV2 cmdDef = 2

type ardrone3MediaRecordStatePictureStateChangedV2 command

type ardrone3MediaRecordStatePictureStateChangedV2Arguments struct {
state uint32
error uint32
}

func (a ardrone3MediaRecordStatePictureStateChangedV2) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordStatePictureStateChangedV2Arguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.error)
offset += 4

return arg
}

var mediaRecordStatePictureStateChangedV2 = ardrone3MediaRecordStatePictureStateChangedV2 {
project: projectardrone3,
class: classMediaRecordState,
cmd: cmdPictureStateChangedV2,
}

// title : Video record state, 
// desc : Video record state., 
// support : 0901:2.0.1;090c;090e, 
// triggered : by [RecordVideo](#1-7-3) or by a change in the video state, 
const cmdVideoStateChangedV2 cmdDef = 3

type ardrone3MediaRecordStateVideoStateChangedV2 command

type ardrone3MediaRecordStateVideoStateChangedV2Arguments struct {
state uint32
error uint32
}

func (a ardrone3MediaRecordStateVideoStateChangedV2) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordStateVideoStateChangedV2Arguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.error)
offset += 4

return arg
}

var mediaRecordStateVideoStateChangedV2 = ardrone3MediaRecordStateVideoStateChangedV2 {
project: projectardrone3,
class: classMediaRecordState,
cmd: cmdVideoStateChangedV2,
}

// title : Video resolution, 
// desc : Video resolution.\n Informs about streaming and recording video resolutions.\n Note that this is only an indication about what the resolution should be. To know the real resolution, you should get it from the frame., 
// support : none, 
// triggered : when the resolution changes., 
const cmdVideoResolutionState cmdDef = 4

type ardrone3MediaRecordStateVideoResolutionState command

type ardrone3MediaRecordStateVideoResolutionStateArguments struct {
streaming uint32
recording uint32
}

func (a ardrone3MediaRecordStateVideoResolutionState) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordStateVideoResolutionStateArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.streaming)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.recording)
offset += 4

return arg
}

var mediaRecordStateVideoResolutionState = ardrone3MediaRecordStateVideoResolutionState {
project: projectardrone3,
class: classMediaRecordState,
cmd: cmdVideoResolutionState,
}

// Events of media recording
const classMediaRecordEvent classDef = 3
// title : Picture taken, 
// desc : Picture taken.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**, 
// support : 0901:2.0.1;090c;090e, 
// triggered : after a [TakePicture](#1-7-2), when the picture has been taken (or it has failed)., 
const cmdPictureEventChanged cmdDef = 0

type ardrone3MediaRecordEventPictureEventChanged command

type ardrone3MediaRecordEventPictureEventChangedArguments struct {
event uint32
error uint32
}

func (a ardrone3MediaRecordEventPictureEventChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordEventPictureEventChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.event)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.error)
offset += 4

return arg
}

var mediaRecordEventPictureEventChanged = ardrone3MediaRecordEventPictureEventChanged {
project: projectardrone3,
class: classMediaRecordEvent,
cmd: cmdPictureEventChanged,
}

// title : Video record notification, 
// desc : Video record notification.\n\n **This event is a notification, you can't retrieve it in the cache of the device controller.**, 
// support : 0901:2.0.1;090c;090e, 
// triggered : by [RecordVideo](#1-7-3) or a change in the video state., 
const cmdVideoEventChanged cmdDef = 1

type ardrone3MediaRecordEventVideoEventChanged command

type ardrone3MediaRecordEventVideoEventChangedArguments struct {
event uint32
error uint32
}

func (a ardrone3MediaRecordEventVideoEventChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaRecordEventVideoEventChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.event)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.error)
offset += 4

return arg
}

var mediaRecordEventVideoEventChanged = ardrone3MediaRecordEventVideoEventChanged {
project: projectardrone3,
class: classMediaRecordEvent,
cmd: cmdVideoEventChanged,
}

// State from drone
const classPilotingState classDef = 4
// title : Flying state, 
// desc : Flying state., 
// support : 0901;090c;090e, 
// triggered : when the flying state changes., 
const cmdFlyingStateChanged cmdDef = 1

type ardrone3PilotingStateFlyingStateChanged command

type ardrone3PilotingStateFlyingStateChangedArguments struct {
state uint32
}

func (a ardrone3PilotingStateFlyingStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateFlyingStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4

return arg
}

var pilotingStateFlyingStateChanged = ardrone3PilotingStateFlyingStateChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdFlyingStateChanged,
}

// title : Alert state, 
// desc : Alert state., 
// support : 0901;090c;090e, 
// triggered : when an alert happens on the drone., 
const cmdAlertStateChanged cmdDef = 2

type ardrone3PilotingStateAlertStateChanged command

type ardrone3PilotingStateAlertStateChangedArguments struct {
state uint32
}

func (a ardrone3PilotingStateAlertStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateAlertStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4

return arg
}

var pilotingStateAlertStateChanged = ardrone3PilotingStateAlertStateChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdAlertStateChanged,
}

// title : Return home state, 
// desc : Return home state.\n Availability is related to gps fix, magnetometer calibration., 
// support : 0901;090c;090e, 
// triggered : by [ReturnHome](#1-0-5) or when the state of the return home changes., 
const cmdNavigateHomeStateChanged cmdDef = 3

type ardrone3PilotingStateNavigateHomeStateChanged command

type ardrone3PilotingStateNavigateHomeStateChangedArguments struct {
state uint32
reason uint32
}

func (a ardrone3PilotingStateNavigateHomeStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateNavigateHomeStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.reason)
offset += 4

return arg
}

var pilotingStateNavigateHomeStateChanged = ardrone3PilotingStateNavigateHomeStateChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdNavigateHomeStateChanged,
}

// title : Drone's position changed, 
// desc : Drone's position changed., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const cmdPositionChanged cmdDef = 4

type ardrone3PilotingStatePositionChanged command

type ardrone3PilotingStatePositionChangedArguments struct {
latitude float64
longitude float64
altitude float64
}

func (a ardrone3PilotingStatePositionChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStatePositionChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8

return arg
}

var pilotingStatePositionChanged = ardrone3PilotingStatePositionChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdPositionChanged,
}

// title : Drone's speed changed, 
// desc : Drone's speed changed.\n Expressed in the NED referential (North-East-Down)., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const cmdSpeedChanged cmdDef = 5

type ardrone3PilotingStateSpeedChanged command

type ardrone3PilotingStateSpeedChangedArguments struct {
speedX float32
speedY float32
speedZ float32
}

func (a ardrone3PilotingStateSpeedChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateSpeedChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.speedX)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.speedY)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.speedZ)
offset += 4

return arg
}

var pilotingStateSpeedChanged = ardrone3PilotingStateSpeedChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdSpeedChanged,
}

// title : Drone's attitude changed, 
// desc : Drone's attitude changed., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const cmdAttitudeChanged cmdDef = 6

type ardrone3PilotingStateAttitudeChanged command

type ardrone3PilotingStateAttitudeChangedArguments struct {
roll float32
pitch float32
yaw float32
}

func (a ardrone3PilotingStateAttitudeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateAttitudeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.roll)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.pitch)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.yaw)
offset += 4

return arg
}

var pilotingStateAttitudeChanged = ardrone3PilotingStateAttitudeChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdAttitudeChanged,
}

// title : Auto takeoff mode, 
// desc : Auto takeoff mode, 
const cmdAutoTakeOffModeChanged cmdDef = 7

type ardrone3PilotingStateAutoTakeOffModeChanged command

type ardrone3PilotingStateAutoTakeOffModeChangedArguments struct {
state uint8
}

func (a ardrone3PilotingStateAutoTakeOffModeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateAutoTakeOffModeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.state)
offset++ 

return arg
}

var pilotingStateAutoTakeOffModeChanged = ardrone3PilotingStateAutoTakeOffModeChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdAutoTakeOffModeChanged,
}

// title : Drone's altitude changed, 
// desc : Drone's altitude changed.\n The altitude reported is the altitude above the take off point.\n To get the altitude above sea level, see [PositionChanged](#1-4-4)., 
// support : 0901;090c;090e, 
// triggered : regularly., 
const cmdAltitudeChanged cmdDef = 8

type ardrone3PilotingStateAltitudeChanged command

type ardrone3PilotingStateAltitudeChangedArguments struct {
altitude float64
}

func (a ardrone3PilotingStateAltitudeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateAltitudeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8

return arg
}

var pilotingStateAltitudeChanged = ardrone3PilotingStateAltitudeChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdAltitudeChanged,
}

// title : Drone's location changed, 
// desc : Drone's location changed.\n This event is meant to replace [PositionChanged](#1-4-4)., 
// support : 0901:4.0.0;090c:4.0.0, 
// triggered : regularly., 
const cmdGpsLocationChanged cmdDef = 9

type ardrone3PilotingStateGpsLocationChanged command

type ardrone3PilotingStateGpsLocationChangedArguments struct {
latitude float64
longitude float64
altitude float64
latitudeaccuracy int8
longitudeaccuracy int8
altitudeaccuracy int8
}

func (a ardrone3PilotingStateGpsLocationChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateGpsLocationChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.latitudeaccuracy)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.longitudeaccuracy)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.altitudeaccuracy)
offset++ 

return arg
}

var pilotingStateGpsLocationChanged = ardrone3PilotingStateGpsLocationChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdGpsLocationChanged,
}

// title : Landing state, 
// desc : Landing state.\n Only available for fixed wings (which have two landing modes)., 
// support : 090e, 
// triggered : when the landing state changes., 
const cmdLandingStateChanged cmdDef = 10

type ardrone3PilotingStateLandingStateChanged command

type ardrone3PilotingStateLandingStateChangedArguments struct {
state uint32
}

func (a ardrone3PilotingStateLandingStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateLandingStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4

return arg
}

var pilotingStateLandingStateChanged = ardrone3PilotingStateLandingStateChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdLandingStateChanged,
}

// title : Drone's air speed changed, 
// desc : Drone's air speed changed\n Expressed in the drone's referential., 
// support : 090e:1.2.0, 
// triggered : regularly., 
const cmdAirSpeedChanged cmdDef = 11

type ardrone3PilotingStateAirSpeedChanged command

type ardrone3PilotingStateAirSpeedChangedArguments struct {
airSpeed float32
}

func (a ardrone3PilotingStateAirSpeedChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateAirSpeedChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.airSpeed)
offset += 4

return arg
}

var pilotingStateAirSpeedChanged = ardrone3PilotingStateAirSpeedChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdAirSpeedChanged,
}

// title : Move to changed, 
// desc : The drone moves or moved to a given location., 
// support : 090c:4.3.0, 
// triggered : by [MoveTo](#1-0-10) or when the drone did reach the given position., 
const cmdmoveToChanged cmdDef = 12

type ardrone3PilotingStatemoveToChanged command

type ardrone3PilotingStatemoveToChangedArguments struct {
latitude float64
longitude float64
altitude float64
orientationmode uint32
heading float32
status uint32
}

func (a ardrone3PilotingStatemoveToChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStatemoveToChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.orientationmode)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.heading)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.status)
offset += 4

return arg
}

var pilotingStatemoveToChanged = ardrone3PilotingStatemoveToChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdmoveToChanged,
}

// title : Motion state, 
// desc : Motion state.\n If [MotionDetection](#1-6-16) is disabled, motion is steady.\n This information is only valid when the drone is not flying., 
// support : 090c:4.3.0, 
// triggered : when the [FlyingState](#1-4-1) is landed and the [MotionDetection](#1-6-16) is enabled and the motion state changes.\n This event is triggered at a filtered rate., 
const cmdMotionState cmdDef = 13

type ardrone3PilotingStateMotionState command

type ardrone3PilotingStateMotionStateArguments struct {
state uint32
}

func (a ardrone3PilotingStateMotionState) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateMotionStateArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4

return arg
}

var pilotingStateMotionState = ardrone3PilotingStateMotionState {
project: projectardrone3,
class: classPilotingState,
cmd: cmdMotionState,
}

// title : Piloted POI state, 
// desc : Piloted POI state., 
// support : 090c:4.3.0, 
// triggered : by [StartPilotedPOI](#1-0-12) or [StopPilotedPOI](#1-0-13) or when piloted POI becomes unavailable., 
const cmdPilotedPOI cmdDef = 14

type ardrone3PilotingStatePilotedPOI command

type ardrone3PilotingStatePilotedPOIArguments struct {
latitude float64
longitude float64
altitude float64
status uint32
}

func (a ardrone3PilotingStatePilotedPOI) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStatePilotedPOIArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.status)
offset += 4

return arg
}

var pilotingStatePilotedPOI = ardrone3PilotingStatePilotedPOI {
project: projectardrone3,
class: classPilotingState,
cmd: cmdPilotedPOI,
}

// title : Return home battery capacity, 
// desc : Battery capacity status to return home., 
// support : 090c:4.3.0, 
// triggered : when the status of the battery capacity to do a return home changes. This means that it is triggered either when the battery level changes, when the distance to the home changes or when the position of the home changes., 
const cmdReturnHomeBatteryCapacity cmdDef = 15

type ardrone3PilotingStateReturnHomeBatteryCapacity command

type ardrone3PilotingStateReturnHomeBatteryCapacityArguments struct {
status uint32
}

func (a ardrone3PilotingStateReturnHomeBatteryCapacity) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateReturnHomeBatteryCapacityArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.status)
offset += 4

return arg
}

var pilotingStateReturnHomeBatteryCapacity = ardrone3PilotingStateReturnHomeBatteryCapacity {
project: projectardrone3,
class: classPilotingState,
cmd: cmdReturnHomeBatteryCapacity,
}

// title : Relative move changed, 
// desc : Relative move changed., 
// triggered : by [MoveRelatively](#1-0-7), or [CancelRelativeMove](#1-0-14) or when the drone's relative move state changes., 
const cmdmoveByChanged cmdDef = 16

type ardrone3PilotingStatemoveByChanged command

type ardrone3PilotingStatemoveByChangedArguments struct {
dXAsked float32
dYAsked float32
dZAsked float32
dPsiAsked float32
dX float32
dY float32
dZ float32
dPsi float32
status uint32
}

func (a ardrone3PilotingStatemoveByChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStatemoveByChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dXAsked)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dYAsked)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dZAsked)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dPsiAsked)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dX)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dY)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dZ)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dPsi)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.status)
offset += 4

return arg
}

var pilotingStatemoveByChanged = ardrone3PilotingStatemoveByChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdmoveByChanged,
}

// title : Hovering warning, 
// desc : Indicate that the drone may have difficulties to maintain a fix position when hovering., 
// support : 0915, 
// triggered : at connection and on changes., 
const cmdHoveringWarning cmdDef = 17

type ardrone3PilotingStateHoveringWarning command

type ardrone3PilotingStateHoveringWarningArguments struct {
nogpstoodark uint8
nogpstoohigh uint8
}

func (a ardrone3PilotingStateHoveringWarning) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateHoveringWarningArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.nogpstoodark)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.nogpstoohigh)
offset++ 

return arg
}

var pilotingStateHoveringWarning = ardrone3PilotingStateHoveringWarning {
project: projectardrone3,
class: classPilotingState,
cmd: cmdHoveringWarning,
}

// title : Landing auto trigger., 
// desc : Forced landing auto trigger information., 
// support : , 
// triggered : at connection, and when forced landing auto trigger information changes, then every seconds while `reason` is different from `none`., 
const cmdForcedLandingAutoTrigger cmdDef = 18

type ardrone3PilotingStateForcedLandingAutoTrigger command

type ardrone3PilotingStateForcedLandingAutoTriggerArguments struct {
reason uint32
delay uint32
}

func (a ardrone3PilotingStateForcedLandingAutoTrigger) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateForcedLandingAutoTriggerArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.reason)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.delay)
offset += 4

return arg
}

var pilotingStateForcedLandingAutoTrigger = ardrone3PilotingStateForcedLandingAutoTrigger {
project: projectardrone3,
class: classPilotingState,
cmd: cmdForcedLandingAutoTrigger,
}

// title : Wind state, 
// desc : Wind state., 
// support : 0914, 
// triggered : at connection and on changes., 
const cmdWindStateChanged cmdDef = 19

type ardrone3PilotingStateWindStateChanged command

type ardrone3PilotingStateWindStateChangedArguments struct {
state uint32
}

func (a ardrone3PilotingStateWindStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingStateWindStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4

return arg
}

var pilotingStateWindStateChanged = ardrone3PilotingStateWindStateChanged {
project: projectardrone3,
class: classPilotingState,
cmd: cmdWindStateChanged,
}

// Events of Piloting
const classPilotingEvent classDef = 34
// title : Relative move ended, 
// desc : Relative move ended.\n Informs about the move that the drone managed to do and why it stopped., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : when the drone reaches its target or when it is interrupted by another [moveBy command](#1-0-7) or when an error occurs., 
const cmdmoveByEnd cmdDef = 0

type ardrone3PilotingEventmoveByEnd command

type ardrone3PilotingEventmoveByEndArguments struct {
dX float32
dY float32
dZ float32
dPsi float32
error uint32
}

func (a ardrone3PilotingEventmoveByEnd) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingEventmoveByEndArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dX)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dY)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dZ)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.dPsi)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.error)
offset += 4

return arg
}

var pilotingEventmoveByEnd = ardrone3PilotingEventmoveByEnd {
project: projectardrone3,
class: classPilotingEvent,
cmd: cmdmoveByEnd,
}

// Network related commands
const classNetwork classDef = 13
// title : Scan wifi network, 
// desc : Scan wifi network to get a list of all networks found by the drone, 
// support : 0901;090c;090e, 
// result : Event [WifiScanResults](#1-14-0) is triggered with all networks found.\n When all networks have been sent, event [WifiScanEnded](#1-14-1) is triggered., 
const cmdWifiScan cmdDef = 0

type ardrone3NetworkWifiScan command

type ardrone3NetworkWifiScanArguments struct {
band uint32
}

func (a ardrone3NetworkWifiScan) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkWifiScanArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.band)
offset += 4

return arg
}

var networkWifiScan = ardrone3NetworkWifiScan {
project: projectardrone3,
class: classNetwork,
cmd: cmdWifiScan,
}

// title : Ask for available wifi channels, 
// desc : Ask for available wifi channels.\n The list of available Wifi channels is related to the country of the drone. You can get this country from the event [CountryChanged](#0-3-6)., 
// support : 0901;090c;090e, 
// result : Event [AvailableWifiChannels](#1-14-2) is triggered with all available channels. When all channels have been sent, event [AvailableWifiChannelsCompleted](#1-14-3) is triggered., 
const cmdWifiAuthChannel cmdDef = 1

type ardrone3NetworkWifiAuthChannel command

type ardrone3NetworkWifiAuthChannelArguments struct {
}

func (a ardrone3NetworkWifiAuthChannel) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkWifiAuthChannelArguments{}
// No arguments to decode here !!

return arg
}

var networkWifiAuthChannel = ardrone3NetworkWifiAuthChannel {
project: projectardrone3,
class: classNetwork,
cmd: cmdWifiAuthChannel,
}

// Network state from Product
const classNetworkState classDef = 14
// title : Wifi scan results, 
// desc : Wifi scan results.\n Please note that the list is not complete until you receive the event [WifiScanEnded](#1-14-1)., 
// support : 0901;090c;090e, 
// triggered : for each wifi network scanned after a [ScanWifi](#1-13-0), 
const cmdWifiScanListChanged cmdDef = 0

type ardrone3NetworkStateWifiScanListChanged command

type ardrone3NetworkStateWifiScanListChangedArguments struct {
ssid string
rssi int16
band uint32
channel uint8
}

func (a ardrone3NetworkStateWifiScanListChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3NetworkStateWifiScanListChangedArguments{}
var offset = 0

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.ssid = string(b[offset:offset+stringEnd])
offset += stringEnd
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.rssi)
offset += 2
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.band)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.channel)
offset++ 

return arg
}

var networkStateWifiScanListChanged = ardrone3NetworkStateWifiScanListChanged {
project: projectardrone3,
class: classNetworkState,
cmd: cmdWifiScanListChanged,
}

// title : Wifi scan ended, 
// desc : Wifi scan ended.\n When receiving this event, the list of [WifiScanResults](#1-14-0) is complete., 
// support : 0901;090c;090e, 
// triggered : after the last [WifiScanResult](#1-14-0) has been sent., 
const cmdAllWifiScanChanged cmdDef = 1

type ardrone3NetworkStateAllWifiScanChanged command

type ardrone3NetworkStateAllWifiScanChangedArguments struct {
}

func (a ardrone3NetworkStateAllWifiScanChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkStateAllWifiScanChangedArguments{}
// No arguments to decode here !!

return arg
}

var networkStateAllWifiScanChanged = ardrone3NetworkStateAllWifiScanChanged {
project: projectardrone3,
class: classNetworkState,
cmd: cmdAllWifiScanChanged,
}

// title : Available wifi channels, 
// desc : Available wifi channels.\n Please note that the list is not complete until you receive the event [AvailableWifiChannelsCompleted](#1-14-3)., 
// support : 0901;090c;090e, 
// triggered : for each available channel after a [GetAvailableWifiChannels](#1-13-1)., 
const cmdWifiAuthChannelListChanged cmdDef = 2

type ardrone3NetworkStateWifiAuthChannelListChanged command

type ardrone3NetworkStateWifiAuthChannelListChangedArguments struct {
band uint32
channel uint8
inorout uint8
}

func (a ardrone3NetworkStateWifiAuthChannelListChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkStateWifiAuthChannelListChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.band)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.channel)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.inorout)
offset++ 

return arg
}

var networkStateWifiAuthChannelListChanged = ardrone3NetworkStateWifiAuthChannelListChanged {
project: projectardrone3,
class: classNetworkState,
cmd: cmdWifiAuthChannelListChanged,
}

// title : Available wifi channels completed, 
// desc : Available wifi channels completed.\n When receiving this event, the list of [AvailableWifiChannels](#1-14-2) is complete., 
// support : 0901;090c;090e, 
// triggered : after the last [AvailableWifiChannel](#1-14-2) has been sent., 
const cmdAllWifiAuthChannelChanged cmdDef = 3

type ardrone3NetworkStateAllWifiAuthChannelChanged command

type ardrone3NetworkStateAllWifiAuthChannelChangedArguments struct {
}

func (a ardrone3NetworkStateAllWifiAuthChannelChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkStateAllWifiAuthChannelChangedArguments{}
// No arguments to decode here !!

return arg
}

var networkStateAllWifiAuthChannelChanged = ardrone3NetworkStateAllWifiAuthChannelChanged {
project: projectardrone3,
class: classNetworkState,
cmd: cmdAllWifiAuthChannelChanged,
}

// Piloting Settings commands
const classPilotingSettings classDef = 2
// title : Set max altitude, 
// desc : Set max altitude.\n The drone will not fly over this max altitude when it is in manual piloting.\n Please note that if you set a max altitude which is below the current drone altitude, the drone will not go to given max altitude.\n You can get the bounds in the event [MaxAltitude](#1-6-0)., 
// support : 0901;090c;090e, 
// result : The max altitude is set.\n Then, event [MaxAltitude](#1-6-0) is triggered., 
const cmdMaxAltitude cmdDef = 0

type ardrone3PilotingSettingsMaxAltitude command

type ardrone3PilotingSettingsMaxAltitudeArguments struct {
current float32
}

func (a ardrone3PilotingSettingsMaxAltitude) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsMaxAltitudeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4

return arg
}

var pilotingSettingsMaxAltitude = ardrone3PilotingSettingsMaxAltitude {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdMaxAltitude,
}

// title : Set max pitch/roll, 
// desc : Set max pitch/roll.\n This represent the max inclination allowed by the drone.\n You can get the bounds with the commands [MaxPitchRoll](#1-6-1)., 
// support : 0901;090c, 
// result : The max pitch/roll is set.\n Then, event [MaxPitchRoll](#1-6-1) is triggered., 
const cmdMaxTilt cmdDef = 1

type ardrone3PilotingSettingsMaxTilt command

type ardrone3PilotingSettingsMaxTiltArguments struct {
current float32
}

func (a ardrone3PilotingSettingsMaxTilt) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsMaxTiltArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4

return arg
}

var pilotingSettingsMaxTilt = ardrone3PilotingSettingsMaxTilt {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdMaxTilt,
}

// title : Set absolut control, 
// desc : Set absolut control., 
const cmdAbsolutControl cmdDef = 2

type ardrone3PilotingSettingsAbsolutControl command

type ardrone3PilotingSettingsAbsolutControlArguments struct {
on uint8
}

func (a ardrone3PilotingSettingsAbsolutControl) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsAbsolutControlArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.on)
offset++ 

return arg
}

var pilotingSettingsAbsolutControl = ardrone3PilotingSettingsAbsolutControl {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdAbsolutControl,
}

// title : Set max distance, 
// desc : Set max distance.\n You can get the bounds from the event [MaxDistance](#1-6-3).\n\n If [Geofence](#1-6-4) is activated, the drone won't fly over the given max distance., 
// support : 0901;090c;090e, 
// result : The max distance is set.\n Then, event [MaxDistance](#1-6-3) is triggered., 
const cmdMaxDistance cmdDef = 3

type ardrone3PilotingSettingsMaxDistance command

type ardrone3PilotingSettingsMaxDistanceArguments struct {
value float32
}

func (a ardrone3PilotingSettingsMaxDistance) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsMaxDistanceArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsMaxDistance = ardrone3PilotingSettingsMaxDistance {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdMaxDistance,
}

// title : Enable geofence, 
// desc : Enable geofence.\n If geofence is enabled, the drone won't fly over the given max distance.\n You can get the max distance from the event [MaxDistance](#1-6-3). \n For copters: the distance is computed from the controller position, if this position is not known, it will use the take off.\n For fixed wings: the distance is computed from the take off position., 
// support : 0901;090c;090e, 
// result : Geofencing is enabled or disabled.\n Then, event [Geofencing](#1-6-4) is triggered., 
const cmdNoFlyOverMaxDistance cmdDef = 4

type ardrone3PilotingSettingsNoFlyOverMaxDistance command

type ardrone3PilotingSettingsNoFlyOverMaxDistanceArguments struct {
shouldNotFlyOver uint8
}

func (a ardrone3PilotingSettingsNoFlyOverMaxDistance) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsNoFlyOverMaxDistanceArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.shouldNotFlyOver)
offset++ 

return arg
}

var pilotingSettingsNoFlyOverMaxDistance = ardrone3PilotingSettingsNoFlyOverMaxDistance {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdNoFlyOverMaxDistance,
}

// title : Set autonomous flight max horizontal speed, 
// desc : Set autonomous flight max horizontal speed.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max horizontal speed is set.\n Then, event [AutonomousFlightMaxHorizontalSpeed](#1-6-5) is triggered., 
const cmdsetAutonomousFlightMaxHorizontalSpeed cmdDef = 5

type ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed command

type ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeedArguments struct {
value float32
}

func (a ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingssetAutonomousFlightMaxHorizontalSpeed = ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalSpeed {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdsetAutonomousFlightMaxHorizontalSpeed,
}

// title : Set autonomous flight max vertical speed, 
// desc : Set autonomous flight max vertical speed.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max vertical speed is set.\n Then, event [AutonomousFlightMaxVerticalSpeed](#1-6-6) is triggered., 
const cmdsetAutonomousFlightMaxVerticalSpeed cmdDef = 6

type ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed command

type ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeedArguments struct {
value float32
}

func (a ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingssetAutonomousFlightMaxVerticalSpeed = ardrone3PilotingSettingssetAutonomousFlightMaxVerticalSpeed {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdsetAutonomousFlightMaxVerticalSpeed,
}

// title : Set autonomous flight max horizontal acceleration, 
// desc : Set autonomous flight max horizontal acceleration.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max horizontal acceleration is set.\n Then, event [AutonomousFlightMaxHorizontalAcceleration](#1-6-7) is triggered., 
const cmdsetAutonomousFlightMaxHorizontalAcceleration cmdDef = 7

type ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration command

type ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAccelerationArguments struct {
value float32
}

func (a ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAccelerationArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingssetAutonomousFlightMaxHorizontalAcceleration = ardrone3PilotingSettingssetAutonomousFlightMaxHorizontalAcceleration {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdsetAutonomousFlightMaxHorizontalAcceleration,
}

// title : Set autonomous flight max vertical acceleration, 
// desc : Set autonomous flight max vertical acceleration.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max vertical acceleration is set.\n Then, event [AutonomousFlightMaxVerticalAcceleration](#1-6-8) is triggered., 
const cmdsetAutonomousFlightMaxVerticalAcceleration cmdDef = 8

type ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration command

type ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAccelerationArguments struct {
value float32
}

func (a ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAccelerationArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingssetAutonomousFlightMaxVerticalAcceleration = ardrone3PilotingSettingssetAutonomousFlightMaxVerticalAcceleration {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdsetAutonomousFlightMaxVerticalAcceleration,
}

// title : Set autonomous flight max rotation speed, 
// desc : Set autonomous flight max rotation speed.\n This will only be used during autonomous flights such as moveBy., 
// support : 0901:3.3.0;090c:3.3.0, 
// result : The max rotation speed is set.\n Then, event [AutonomousFlightMaxRotationSpeed](#1-6-9) is triggered., 
const cmdsetAutonomousFlightMaxRotationSpeed cmdDef = 9

type ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed command

type ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeedArguments struct {
value float32
}

func (a ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingssetAutonomousFlightMaxRotationSpeed = ardrone3PilotingSettingssetAutonomousFlightMaxRotationSpeed {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdsetAutonomousFlightMaxRotationSpeed,
}

// title : Set banked turn mode, 
// desc : Set banked turn mode.\n When banked turn mode is enabled, the drone will use yaw values from the piloting command to infer with roll and pitch on the drone when its horizontal speed is not null., 
// support : 0901:3.2.0;090c:3.2.0, 
// result : The banked turn mode is enabled or disabled.\n Then, event [BankedTurnMode](#1-6-10) is triggered., 
const cmdBankedTurn cmdDef = 10

type ardrone3PilotingSettingsBankedTurn command

type ardrone3PilotingSettingsBankedTurnArguments struct {
value uint8
}

func (a ardrone3PilotingSettingsBankedTurn) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsBankedTurnArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.value)
offset++ 

return arg
}

var pilotingSettingsBankedTurn = ardrone3PilotingSettingsBankedTurn {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdBankedTurn,
}

// title : Set minimum altitude, 
// desc : Set minimum altitude.\n Only available for fixed wings., 
// support : 090e, 
// result : The minimum altitude is set.\n Then, event [MinimumAltitude](#1-6-11) is triggered., 
const cmdMinAltitude cmdDef = 11

type ardrone3PilotingSettingsMinAltitude command

type ardrone3PilotingSettingsMinAltitudeArguments struct {
current float32
}

func (a ardrone3PilotingSettingsMinAltitude) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsMinAltitudeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4

return arg
}

var pilotingSettingsMinAltitude = ardrone3PilotingSettingsMinAltitude {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdMinAltitude,
}

// title : Set default circling direction, 
// desc : Set default circling direction. This direction will be used when the drone use an automatic circling or when [CIRCLE](#1-0-9) is sent with direction *default*.\n Only available for fixed wings., 
// support : 090e, 
// result : The circling direction is set.\n Then, event [DefaultCirclingDirection](#1-6-12) is triggered., 
const cmdCirclingDirection cmdDef = 12

type ardrone3PilotingSettingsCirclingDirection command

type ardrone3PilotingSettingsCirclingDirectionArguments struct {
value uint32
}

func (a ardrone3PilotingSettingsCirclingDirection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsCirclingDirectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsCirclingDirection = ardrone3PilotingSettingsCirclingDirection {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdCirclingDirection,
}

// title : Set circling radius, 
// desc : Set circling radius.\n Only available for fixed wings., 
// support : none, 
// result : The circling radius is set.\n Then, event [CirclingRadius](#1-6-13) is triggered., 
const cmdCirclingRadius cmdDef = 13

type ardrone3PilotingSettingsCirclingRadius command

type ardrone3PilotingSettingsCirclingRadiusArguments struct {
value uint16
}

func (a ardrone3PilotingSettingsCirclingRadius) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsCirclingRadiusArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.value)
offset += 2

return arg
}

var pilotingSettingsCirclingRadius = ardrone3PilotingSettingsCirclingRadius {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdCirclingRadius,
}

// title : Set min circling altitude, 
// desc : Set min circling altitude (not used during take off).\n Only available for fixed wings., 
// support : 090e, 
// result : The circling altitude is set.\n Then, event [CirclingAltitude](#1-6-14) is triggered., 
const cmdCirclingAltitude cmdDef = 14

type ardrone3PilotingSettingsCirclingAltitude command

type ardrone3PilotingSettingsCirclingAltitudeArguments struct {
value uint16
}

func (a ardrone3PilotingSettingsCirclingAltitude) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsCirclingAltitudeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.value)
offset += 2

return arg
}

var pilotingSettingsCirclingAltitude = ardrone3PilotingSettingsCirclingAltitude {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdCirclingAltitude,
}

// title : Set pitch mode, 
// desc : Set pitch mode.\n Only available for fixed wings., 
// support : 090e, 
// result : The pitch mode is set.\n Then, event [PitchMode](#1-6-15) is triggered., 
const cmdPitchMode cmdDef = 15

type ardrone3PilotingSettingsPitchMode command

type ardrone3PilotingSettingsPitchModeArguments struct {
value uint32
}

func (a ardrone3PilotingSettingsPitchMode) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsPitchModeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsPitchMode = ardrone3PilotingSettingsPitchMode {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdPitchMode,
}

// title : Enable/disable the motion detection, 
// desc : Enable/disable the motion detection.\n If the motion detection is enabled, the drone will send its [MotionState](#1-4-13) when its [FlyingState](#1-4-1) is landed. If the motion detection is disabled, [MotionState](#1-4-13) is steady., 
// support : 090c:4.3.0, 
// result : The motion detection is enabled or disabled.\n Then, event [MotionDetection](#1-6-16) is triggered. After that, if enabled and [FlyingState](#1-4-1) is landed, the [MotionState](#1-4-13) is triggered upon changes., 
const cmdSetMotionDetectionMode cmdDef = 16

type ardrone3PilotingSettingsSetMotionDetectionMode command

type ardrone3PilotingSettingsSetMotionDetectionModeArguments struct {
enable uint8
}

func (a ardrone3PilotingSettingsSetMotionDetectionMode) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsSetMotionDetectionModeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.enable)
offset++ 

return arg
}

var pilotingSettingsSetMotionDetectionMode = ardrone3PilotingSettingsSetMotionDetectionMode {
project: projectardrone3,
class: classPilotingSettings,
cmd: cmdSetMotionDetectionMode,
}

// Piloting Settings state from product
const classPilotingSettingsState classDef = 6
// title : Max altitude, 
// desc : Max altitude.\n The drone will not fly higher than this altitude (above take off point)., 
// support : 0901;090c;090e, 
// triggered : by [SetMaxAltitude](#1-2-0)., 
const cmdMaxAltitudeChanged cmdDef = 0

type ardrone3PilotingSettingsStateMaxAltitudeChanged command

type ardrone3PilotingSettingsStateMaxAltitudeChangedArguments struct {
current float32
min float32
max float32
}

func (a ardrone3PilotingSettingsStateMaxAltitudeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateMaxAltitudeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var pilotingSettingsStateMaxAltitudeChanged = ardrone3PilotingSettingsStateMaxAltitudeChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdMaxAltitudeChanged,
}

// title : Max pitch/roll, 
// desc : Max pitch/roll.\n The drone will not fly higher than this altitude (above take off point)., 
// support : 0901;090c, 
// triggered : by [SetMaxAltitude](#1-2-0)., 
const cmdMaxTiltChanged cmdDef = 1

type ardrone3PilotingSettingsStateMaxTiltChanged command

type ardrone3PilotingSettingsStateMaxTiltChangedArguments struct {
current float32
min float32
max float32
}

func (a ardrone3PilotingSettingsStateMaxTiltChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateMaxTiltChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var pilotingSettingsStateMaxTiltChanged = ardrone3PilotingSettingsStateMaxTiltChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdMaxTiltChanged,
}

// title : Absolut control, 
// desc : Absolut control., 
const cmdAbsolutControlChanged cmdDef = 2

type ardrone3PilotingSettingsStateAbsolutControlChanged command

type ardrone3PilotingSettingsStateAbsolutControlChangedArguments struct {
on uint8
}

func (a ardrone3PilotingSettingsStateAbsolutControlChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateAbsolutControlChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.on)
offset++ 

return arg
}

var pilotingSettingsStateAbsolutControlChanged = ardrone3PilotingSettingsStateAbsolutControlChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdAbsolutControlChanged,
}

// title : Max distance, 
// desc : Max distance., 
// support : 0901;090c;090e, 
// triggered : by [SetMaxDistance](#1-2-3)., 
const cmdMaxDistanceChanged cmdDef = 3

type ardrone3PilotingSettingsStateMaxDistanceChanged command

type ardrone3PilotingSettingsStateMaxDistanceChangedArguments struct {
current float32
min float32
max float32
}

func (a ardrone3PilotingSettingsStateMaxDistanceChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateMaxDistanceChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var pilotingSettingsStateMaxDistanceChanged = ardrone3PilotingSettingsStateMaxDistanceChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdMaxDistanceChanged,
}

// title : Geofencing, 
// desc : Geofencing.\n If set, the drone won't fly over the [MaxDistance](#1-6-3)., 
// support : 0901;090c;090e, 
// triggered : by [EnableGeofence](#1-2-4)., 
const cmdNoFlyOverMaxDistanceChanged cmdDef = 4

type ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged command

type ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChangedArguments struct {
shouldNotFlyOver uint8
}

func (a ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.shouldNotFlyOver)
offset++ 

return arg
}

var pilotingSettingsStateNoFlyOverMaxDistanceChanged = ardrone3PilotingSettingsStateNoFlyOverMaxDistanceChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdNoFlyOverMaxDistanceChanged,
}

// title : Autonomous flight max horizontal speed, 
// desc : Autonomous flight max horizontal speed., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxHorizontalSpeed](#1-2-5)., 
const cmdAutonomousFlightMaxHorizontalSpeed cmdDef = 5

type ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed command

type ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeedArguments struct {
value float32
}

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsStateAutonomousFlightMaxHorizontalSpeed = ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalSpeed {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdAutonomousFlightMaxHorizontalSpeed,
}

// title : Autonomous flight max vertical speed, 
// desc : Autonomous flight max vertical speed., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxVerticalSpeed](#1-2-6)., 
const cmdAutonomousFlightMaxVerticalSpeed cmdDef = 6

type ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed command

type ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeedArguments struct {
value float32
}

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsStateAutonomousFlightMaxVerticalSpeed = ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalSpeed {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdAutonomousFlightMaxVerticalSpeed,
}

// title : Autonomous flight max horizontal acceleration, 
// desc : Autonomous flight max horizontal acceleration., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxHorizontalAcceleration](#1-2-7)., 
const cmdAutonomousFlightMaxHorizontalAcceleration cmdDef = 7

type ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration command

type ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAccelerationArguments struct {
value float32
}

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAccelerationArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration = ardrone3PilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdAutonomousFlightMaxHorizontalAcceleration,
}

// title : Autonomous flight max vertical acceleration, 
// desc : Autonomous flight max vertical acceleration., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxVerticalAcceleration](#1-2-8)., 
const cmdAutonomousFlightMaxVerticalAcceleration cmdDef = 8

type ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration command

type ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAccelerationArguments struct {
value float32
}

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAccelerationArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsStateAutonomousFlightMaxVerticalAcceleration = ardrone3PilotingSettingsStateAutonomousFlightMaxVerticalAcceleration {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdAutonomousFlightMaxVerticalAcceleration,
}

// title : Autonomous flight max rotation speed, 
// desc : Autonomous flight max rotation speed., 
// support : 0901:3.3.0;090c:3.3.0, 
// triggered : by [SetAutonomousFlightMaxRotationSpeed](#1-2-9)., 
const cmdAutonomousFlightMaxRotationSpeed cmdDef = 9

type ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed command

type ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeedArguments struct {
value float32
}

func (a ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsStateAutonomousFlightMaxRotationSpeed = ardrone3PilotingSettingsStateAutonomousFlightMaxRotationSpeed {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdAutonomousFlightMaxRotationSpeed,
}

// title : Banked Turn mode, 
// desc : Banked Turn mode.\n If banked turn mode is enabled, the drone will use yaw values from the piloting command to infer with roll and pitch on the drone when its horizontal speed is not null., 
// support : 0901:3.2.0;090c:3.2.0, 
// triggered : by [SetBankedTurnMode](#1-2-10)., 
const cmdBankedTurnChanged cmdDef = 10

type ardrone3PilotingSettingsStateBankedTurnChanged command

type ardrone3PilotingSettingsStateBankedTurnChangedArguments struct {
state uint8
}

func (a ardrone3PilotingSettingsStateBankedTurnChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateBankedTurnChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.state)
offset++ 

return arg
}

var pilotingSettingsStateBankedTurnChanged = ardrone3PilotingSettingsStateBankedTurnChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdBankedTurnChanged,
}

// title : Min altitude, 
// desc : Min altitude.\n Only sent by fixed wings., 
// support : 090e, 
// triggered : by [SetMinAltitude](#1-2-11)., 
const cmdMinAltitudeChanged cmdDef = 11

type ardrone3PilotingSettingsStateMinAltitudeChanged command

type ardrone3PilotingSettingsStateMinAltitudeChangedArguments struct {
current float32
min float32
max float32
}

func (a ardrone3PilotingSettingsStateMinAltitudeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateMinAltitudeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var pilotingSettingsStateMinAltitudeChanged = ardrone3PilotingSettingsStateMinAltitudeChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdMinAltitudeChanged,
}

// title : Circling direction, 
// desc : Circling direction.\n Only sent by fixed wings., 
// support : 090e, 
// triggered : by [SetCirclingDirection](#1-2-12)., 
const cmdCirclingDirectionChanged cmdDef = 12

type ardrone3PilotingSettingsStateCirclingDirectionChanged command

type ardrone3PilotingSettingsStateCirclingDirectionChangedArguments struct {
value uint32
}

func (a ardrone3PilotingSettingsStateCirclingDirectionChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateCirclingDirectionChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsStateCirclingDirectionChanged = ardrone3PilotingSettingsStateCirclingDirectionChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdCirclingDirectionChanged,
}

// title : Circling radius, 
// desc : Circling radius.\n Only sent by fixed wings., 
// support : none, 
// triggered : by [SetCirclingRadius](#1-2-13)., 
const cmdCirclingRadiusChanged cmdDef = 13

type ardrone3PilotingSettingsStateCirclingRadiusChanged command

type ardrone3PilotingSettingsStateCirclingRadiusChangedArguments struct {
current uint16
min uint16
max uint16
}

func (a ardrone3PilotingSettingsStateCirclingRadiusChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateCirclingRadiusChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.current)
offset += 2
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.min)
offset += 2
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.max)
offset += 2

return arg
}

var pilotingSettingsStateCirclingRadiusChanged = ardrone3PilotingSettingsStateCirclingRadiusChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdCirclingRadiusChanged,
}

// title : Circling altitude, 
// desc : Circling altitude.\n Bounds will be automatically adjusted according to the [MaxAltitude](#1-6-0).\n Only sent by fixed wings., 
// support : 090e, 
// triggered : by [SetCirclingRadius](#1-2-14) or when bounds change due to [SetMaxAltitude](#1-2-0)., 
const cmdCirclingAltitudeChanged cmdDef = 14

type ardrone3PilotingSettingsStateCirclingAltitudeChanged command

type ardrone3PilotingSettingsStateCirclingAltitudeChangedArguments struct {
current uint16
min uint16
max uint16
}

func (a ardrone3PilotingSettingsStateCirclingAltitudeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateCirclingAltitudeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.current)
offset += 2
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.min)
offset += 2
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.max)
offset += 2

return arg
}

var pilotingSettingsStateCirclingAltitudeChanged = ardrone3PilotingSettingsStateCirclingAltitudeChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdCirclingAltitudeChanged,
}

// title : Pitch mode, 
// desc : Pitch mode., 
// support : 090e, 
// triggered : by [SetPitchMode](#1-2-15)., 
const cmdPitchModeChanged cmdDef = 15

type ardrone3PilotingSettingsStatePitchModeChanged command

type ardrone3PilotingSettingsStatePitchModeChangedArguments struct {
value uint32
}

func (a ardrone3PilotingSettingsStatePitchModeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStatePitchModeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pilotingSettingsStatePitchModeChanged = ardrone3PilotingSettingsStatePitchModeChanged {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdPitchModeChanged,
}

// title : State of the motion detection, 
// desc : State of the motion detection., 
// support : 090c:4.3.0, 
// triggered : by [SetMotionDetectionMode](#1-2-16), 
const cmdMotionDetection cmdDef = 16

type ardrone3PilotingSettingsStateMotionDetection command

type ardrone3PilotingSettingsStateMotionDetectionArguments struct {
enabled uint8
}

func (a ardrone3PilotingSettingsStateMotionDetection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PilotingSettingsStateMotionDetectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.enabled)
offset++ 

return arg
}

var pilotingSettingsStateMotionDetection = ardrone3PilotingSettingsStateMotionDetection {
project: projectardrone3,
class: classPilotingSettingsState,
cmd: cmdMotionDetection,
}

// Speed Settings commands
const classSpeedSettings classDef = 11
// title : Set max vertical speed, 
// desc : Set max vertical speed., 
// support : 0901;090c, 
// result : The max vertical speed is set.\n Then, event [MaxVerticalSpeed](#1-12-0) is triggered., 
const cmdMaxVerticalSpeed cmdDef = 0

type ardrone3SpeedSettingsMaxVerticalSpeed command

type ardrone3SpeedSettingsMaxVerticalSpeedArguments struct {
current float32
}

func (a ardrone3SpeedSettingsMaxVerticalSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsMaxVerticalSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4

return arg
}

var speedSettingsMaxVerticalSpeed = ardrone3SpeedSettingsMaxVerticalSpeed {
project: projectardrone3,
class: classSpeedSettings,
cmd: cmdMaxVerticalSpeed,
}

// title : Set max rotation speed, 
// desc : Set max rotation speed., 
// support : 0901;090c, 
// result : The max rotation speed is set.\n Then, event [MaxRotationSpeed](#1-12-1) is triggered., 
const cmdMaxRotationSpeed cmdDef = 1

type ardrone3SpeedSettingsMaxRotationSpeed command

type ardrone3SpeedSettingsMaxRotationSpeedArguments struct {
current float32
}

func (a ardrone3SpeedSettingsMaxRotationSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsMaxRotationSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4

return arg
}

var speedSettingsMaxRotationSpeed = ardrone3SpeedSettingsMaxRotationSpeed {
project: projectardrone3,
class: classSpeedSettings,
cmd: cmdMaxRotationSpeed,
}

// title : Set the presence of hull protection, 
// desc : Set the presence of hull protection., 
// support : 0901;090c, 
// result : The drone knows that it has a hull protection.\n Then, event [HullProtection](#1-12-2) is triggered., 
const cmdHullProtection cmdDef = 2

type ardrone3SpeedSettingsHullProtection command

type ardrone3SpeedSettingsHullProtectionArguments struct {
present uint8
}

func (a ardrone3SpeedSettingsHullProtection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsHullProtectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.present)
offset++ 

return arg
}

var speedSettingsHullProtection = ardrone3SpeedSettingsHullProtection {
project: projectardrone3,
class: classSpeedSettings,
cmd: cmdHullProtection,
}

// title : Set outdoor mode, 
// desc : Set outdoor mode., 
const cmdOutdoor cmdDef = 3

type ardrone3SpeedSettingsOutdoor command

type ardrone3SpeedSettingsOutdoorArguments struct {
outdoor uint8
}

func (a ardrone3SpeedSettingsOutdoor) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsOutdoorArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.outdoor)
offset++ 

return arg
}

var speedSettingsOutdoor = ardrone3SpeedSettingsOutdoor {
project: projectardrone3,
class: classSpeedSettings,
cmd: cmdOutdoor,
}

// title : Set max pitch/roll rotation speed, 
// desc : Set max pitch/roll rotation speed., 
// support : 0901;090c, 
// result : The max pitch/roll rotation speed is set.\n Then, event [MaxPitchRollRotationSpeed](#1-12-4) is triggered., 
const cmdMaxPitchRollRotationSpeed cmdDef = 4

type ardrone3SpeedSettingsMaxPitchRollRotationSpeed command

type ardrone3SpeedSettingsMaxPitchRollRotationSpeedArguments struct {
current float32
}

func (a ardrone3SpeedSettingsMaxPitchRollRotationSpeed) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsMaxPitchRollRotationSpeedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4

return arg
}

var speedSettingsMaxPitchRollRotationSpeed = ardrone3SpeedSettingsMaxPitchRollRotationSpeed {
project: projectardrone3,
class: classSpeedSettings,
cmd: cmdMaxPitchRollRotationSpeed,
}

// Speed Settings state from product
const classSpeedSettingsState classDef = 12
// title : Max vertical speed, 
// desc : Max vertical speed., 
// support : 0901;090c, 
// triggered : by [SetMaxVerticalSpeed](#1-11-0)., 
const cmdMaxVerticalSpeedChanged cmdDef = 0

type ardrone3SpeedSettingsStateMaxVerticalSpeedChanged command

type ardrone3SpeedSettingsStateMaxVerticalSpeedChangedArguments struct {
current float32
min float32
max float32
}

func (a ardrone3SpeedSettingsStateMaxVerticalSpeedChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsStateMaxVerticalSpeedChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var speedSettingsStateMaxVerticalSpeedChanged = ardrone3SpeedSettingsStateMaxVerticalSpeedChanged {
project: projectardrone3,
class: classSpeedSettingsState,
cmd: cmdMaxVerticalSpeedChanged,
}

// title : Max rotation speed, 
// desc : Max rotation speed., 
// support : 0901;090c, 
// triggered : by [SetMaxRotationSpeed](#1-11-1)., 
const cmdMaxRotationSpeedChanged cmdDef = 1

type ardrone3SpeedSettingsStateMaxRotationSpeedChanged command

type ardrone3SpeedSettingsStateMaxRotationSpeedChangedArguments struct {
current float32
min float32
max float32
}

func (a ardrone3SpeedSettingsStateMaxRotationSpeedChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsStateMaxRotationSpeedChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var speedSettingsStateMaxRotationSpeedChanged = ardrone3SpeedSettingsStateMaxRotationSpeedChanged {
project: projectardrone3,
class: classSpeedSettingsState,
cmd: cmdMaxRotationSpeedChanged,
}

// title : Presence of hull protection, 
// desc : Presence of hull protection., 
// support : 0901;090c, 
// triggered : by [SetHullProtectionPresence](#1-11-2)., 
const cmdHullProtectionChanged cmdDef = 2

type ardrone3SpeedSettingsStateHullProtectionChanged command

type ardrone3SpeedSettingsStateHullProtectionChangedArguments struct {
present uint8
}

func (a ardrone3SpeedSettingsStateHullProtectionChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsStateHullProtectionChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.present)
offset++ 

return arg
}

var speedSettingsStateHullProtectionChanged = ardrone3SpeedSettingsStateHullProtectionChanged {
project: projectardrone3,
class: classSpeedSettingsState,
cmd: cmdHullProtectionChanged,
}

// title : Outdoor mode, 
// desc : Outdoor mode., 
const cmdOutdoorChanged cmdDef = 3

type ardrone3SpeedSettingsStateOutdoorChanged command

type ardrone3SpeedSettingsStateOutdoorChangedArguments struct {
outdoor uint8
}

func (a ardrone3SpeedSettingsStateOutdoorChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsStateOutdoorChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.outdoor)
offset++ 

return arg
}

var speedSettingsStateOutdoorChanged = ardrone3SpeedSettingsStateOutdoorChanged {
project: projectardrone3,
class: classSpeedSettingsState,
cmd: cmdOutdoorChanged,
}

// title : Max pitch/roll rotation speed, 
// desc : Max pitch/roll rotation speed., 
// support : 0901;090c, 
// triggered : by [SetMaxPitchRollRotationSpeed](#1-11-4)., 
const cmdMaxPitchRollRotationSpeedChanged cmdDef = 4

type ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged command

type ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChangedArguments struct {
current float32
min float32
max float32
}

func (a ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.current)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var speedSettingsStateMaxPitchRollRotationSpeedChanged = ardrone3SpeedSettingsStateMaxPitchRollRotationSpeedChanged {
project: projectardrone3,
class: classSpeedSettingsState,
cmd: cmdMaxPitchRollRotationSpeedChanged,
}

// Network settings commands
const classNetworkSettings classDef = 9
// title : Select Wifi, 
// desc : Select or auto-select channel of choosen band., 
// support : 0901;090c;090e, 
// result : The wifi channel changes according to given parameters. Watch out, a disconnection might appear.\n Then, event [WifiSelection](#1-10-0) is triggered., 
const cmdWifiSelection cmdDef = 0

type ardrone3NetworkSettingsWifiSelection command

type ardrone3NetworkSettingsWifiSelectionArguments struct {
typeX uint32
band uint32
channel uint8
}

func (a ardrone3NetworkSettingsWifiSelection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkSettingsWifiSelectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.band)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.channel)
offset++ 

return arg
}

var networkSettingsWifiSelection = ardrone3NetworkSettingsWifiSelection {
project: projectardrone3,
class: classNetworkSettings,
cmd: cmdWifiSelection,
}

// title : Set wifi security type, 
// desc : Set wifi security type.\n The security will be changed on the next restart, 
// support : 0901;090c;090e, 
// result : The wifi security is set (but not applied until next restart).\n Then, event [WifiSecurityType](#1-10-2) is triggered., 
const cmdwifiSecurity cmdDef = 1

type ardrone3NetworkSettingswifiSecurity command

type ardrone3NetworkSettingswifiSecurityArguments struct {
typeX uint32
key string
keyType uint32
}

func (a ardrone3NetworkSettingswifiSecurity) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3NetworkSettingswifiSecurityArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.key = string(b[offset:offset+stringEnd])
offset += stringEnd
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.keyType)
offset += 4

return arg
}

var networkSettingswifiSecurity = ardrone3NetworkSettingswifiSecurity {
project: projectardrone3,
class: classNetworkSettings,
cmd: cmdwifiSecurity,
}

// Network settings state from product
const classNetworkSettingsState classDef = 10
// title : Wifi selection, 
// desc : Wifi selection., 
// support : 0901;090c;090e, 
// triggered : by [SelectWifi](#1-9-0)., 
const cmdWifiSelectionChanged cmdDef = 0

type ardrone3NetworkSettingsStateWifiSelectionChanged command

type ardrone3NetworkSettingsStateWifiSelectionChangedArguments struct {
typeX uint32
band uint32
channel uint8
}

func (a ardrone3NetworkSettingsStateWifiSelectionChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkSettingsStateWifiSelectionChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.band)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.channel)
offset++ 

return arg
}

var networkSettingsStateWifiSelectionChanged = ardrone3NetworkSettingsStateWifiSelectionChanged {
project: projectardrone3,
class: classNetworkSettingsState,
cmd: cmdWifiSelectionChanged,
}

// title : Wifi security type, 
// desc : Wifi security type., 
const cmdwifiSecurityChanged cmdDef = 1

type ardrone3NetworkSettingsStatewifiSecurityChanged command

type ardrone3NetworkSettingsStatewifiSecurityChangedArguments struct {
typeX uint32
}

func (a ardrone3NetworkSettingsStatewifiSecurityChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3NetworkSettingsStatewifiSecurityChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var networkSettingsStatewifiSecurityChanged = ardrone3NetworkSettingsStatewifiSecurityChanged {
project: projectardrone3,
class: classNetworkSettingsState,
cmd: cmdwifiSecurityChanged,
}

// title : Wifi security type, 
// desc : Wifi security type., 
// support : 0901;090c;090e, 
// triggered : by [SetWifiSecurityType](#1-9-1)., 
const cmdwifiSecurityDUPLICATE cmdDef = 2

type ardrone3NetworkSettingsStatewifiSecurity command

type ardrone3NetworkSettingsStatewifiSecurityArguments struct {
typeX uint32
key string
keyType uint32
}

func (a ardrone3NetworkSettingsStatewifiSecurity) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3NetworkSettingsStatewifiSecurityArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.key = string(b[offset:offset+stringEnd])
offset += stringEnd
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.keyType)
offset += 4

return arg
}

var networkSettingsStatewifiSecurity = ardrone3NetworkSettingsStatewifiSecurity {
project: projectardrone3,
class: classNetworkSettingsState,
cmd: cmdwifiSecurity,
}

// Settings state from product
const classSettingsState classDef = 16
// title : Motor version, 
// desc : Motor version., 
const cmdProductMotorVersionListChanged cmdDef = 0

type ardrone3SettingsStateProductMotorVersionListChanged command

type ardrone3SettingsStateProductMotorVersionListChangedArguments struct {
motornumber uint8
typeX string
software string
hardware string
}

func (a ardrone3SettingsStateProductMotorVersionListChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3SettingsStateProductMotorVersionListChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.motornumber)
offset++ 

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.typeX = string(b[offset:offset+stringEnd])
offset += stringEnd

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.software = string(b[offset:offset+stringEnd])
offset += stringEnd

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.hardware = string(b[offset:offset+stringEnd])
offset += stringEnd

return arg
}

var settingsStateProductMotorVersionListChanged = ardrone3SettingsStateProductMotorVersionListChanged {
project: projectardrone3,
class: classSettingsState,
cmd: cmdProductMotorVersionListChanged,
}

// title : GPS version, 
// desc : GPS version., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const cmdProductGPSVersionChanged cmdDef = 1

type ardrone3SettingsStateProductGPSVersionChanged command

type ardrone3SettingsStateProductGPSVersionChangedArguments struct {
software string
hardware string
}

func (a ardrone3SettingsStateProductGPSVersionChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3SettingsStateProductGPSVersionChangedArguments{}
var offset = 0

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.software = string(b[offset:offset+stringEnd])
offset += stringEnd

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.hardware = string(b[offset:offset+stringEnd])
offset += stringEnd

return arg
}

var settingsStateProductGPSVersionChanged = ardrone3SettingsStateProductGPSVersionChanged {
project: projectardrone3,
class: classSettingsState,
cmd: cmdProductGPSVersionChanged,
}

// title : Motor error, 
// desc : Motor error.\n This event is sent back to *noError* as soon as the motor error disappear. To get the last motor error, see [LastMotorError](#1-16-5), 
// support : 0901;090c;090e, 
// triggered : when a motor error occurs., 
const cmdMotorErrorStateChanged cmdDef = 2

type ardrone3SettingsStateMotorErrorStateChanged command

type ardrone3SettingsStateMotorErrorStateChangedArguments struct {
motorIds uint8
motorError uint32
}

func (a ardrone3SettingsStateMotorErrorStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SettingsStateMotorErrorStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.motorIds)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.motorError)
offset += 4

return arg
}

var settingsStateMotorErrorStateChanged = ardrone3SettingsStateMotorErrorStateChanged {
project: projectardrone3,
class: classSettingsState,
cmd: cmdMotorErrorStateChanged,
}

// title : Motor version, 
// desc : Motor version., 
const cmdMotorSoftwareVersionChanged cmdDef = 3

type ardrone3SettingsStateMotorSoftwareVersionChanged command

type ardrone3SettingsStateMotorSoftwareVersionChangedArguments struct {
version string
}

func (a ardrone3SettingsStateMotorSoftwareVersionChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3SettingsStateMotorSoftwareVersionChangedArguments{}
var offset = 0

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.version = string(b[offset:offset+stringEnd])
offset += stringEnd

return arg
}

var settingsStateMotorSoftwareVersionChanged = ardrone3SettingsStateMotorSoftwareVersionChanged {
project: projectardrone3,
class: classSettingsState,
cmd: cmdMotorSoftwareVersionChanged,
}

// title : Motor flight status, 
// desc : Motor flight status., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const cmdMotorFlightsStatusChanged cmdDef = 4

type ardrone3SettingsStateMotorFlightsStatusChanged command

type ardrone3SettingsStateMotorFlightsStatusChangedArguments struct {
nbFlights uint16
lastFlightDuration uint16
totalFlightDuration uint32
}

func (a ardrone3SettingsStateMotorFlightsStatusChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SettingsStateMotorFlightsStatusChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.nbFlights)
offset += 2
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.lastFlightDuration)
offset += 2
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.totalFlightDuration)
offset += 4

return arg
}

var settingsStateMotorFlightsStatusChanged = ardrone3SettingsStateMotorFlightsStatusChanged {
project: projectardrone3,
class: classSettingsState,
cmd: cmdMotorFlightsStatusChanged,
}

// title : Last motor error, 
// desc : Last motor error.\n This is a reminder of the last error. To know if a motor error is currently happening, see [MotorError](#1-16-2)., 
// support : 0901;090c;090e, 
// triggered : at connection and when an error occurs., 
const cmdMotorErrorLastErrorChanged cmdDef = 5

type ardrone3SettingsStateMotorErrorLastErrorChanged command

type ardrone3SettingsStateMotorErrorLastErrorChangedArguments struct {
motorError uint32
}

func (a ardrone3SettingsStateMotorErrorLastErrorChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SettingsStateMotorErrorLastErrorChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.motorError)
offset += 4

return arg
}

var settingsStateMotorErrorLastErrorChanged = ardrone3SettingsStateMotorErrorLastErrorChanged {
project: projectardrone3,
class: classSettingsState,
cmd: cmdMotorErrorLastErrorChanged,
}

// title : P7ID, 
// desc : P7ID., 
const cmdP7ID cmdDef = 6

type ardrone3SettingsStateP7ID command

type ardrone3SettingsStateP7IDArguments struct {
serialID string
}

func (a ardrone3SettingsStateP7ID) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3SettingsStateP7IDArguments{}
var offset = 0

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.serialID = string(b[offset:offset+stringEnd])
offset += stringEnd

return arg
}

var settingsStateP7ID = ardrone3SettingsStateP7ID {
project: projectardrone3,
class: classSettingsState,
cmd: cmdP7ID,
}

const cmdCPUID cmdDef = 7

type ardrone3SettingsStateCPUID command

type ardrone3SettingsStateCPUIDArguments struct {
id string
}

func (a ardrone3SettingsStateCPUID) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3SettingsStateCPUIDArguments{}
var offset = 0

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.id = string(b[offset:offset+stringEnd])
offset += stringEnd

return arg
}

var settingsStateCPUID = ardrone3SettingsStateCPUID {
project: projectardrone3,
class: classSettingsState,
cmd: cmdCPUID,
}

// Photo settings chosen by the user
const classPictureSettings classDef = 19
// title : Set picture format, 
// desc : Set picture format.\n Please note that the time required to take the picture is highly related to this format.\n Also, please note that if your picture format is different from snapshot, picture taking will stop video recording (it will restart after the picture has been taken)., 
// support : 0901;090c;090e, 
// result : The picture format is set.\n Then, event [PictureFormat](#1-20-0) is triggered., 
const cmdPictureFormatSelection cmdDef = 0

type ardrone3PictureSettingsPictureFormatSelection command

type ardrone3PictureSettingsPictureFormatSelectionArguments struct {
typeX uint32
}

func (a ardrone3PictureSettingsPictureFormatSelection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsPictureFormatSelectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var pictureSettingsPictureFormatSelection = ardrone3PictureSettingsPictureFormatSelection {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdPictureFormatSelection,
}

// title : Set White Balance mode, 
// desc : Set White Balance mode., 
// support : 0901;090c;090e, 
// result : The white balance mode is set.\n Then, event [WhiteBalanceMode](#1-20-1) is triggered., 
const cmdAutoWhiteBalanceSelection cmdDef = 1

type ardrone3PictureSettingsAutoWhiteBalanceSelection command

type ardrone3PictureSettingsAutoWhiteBalanceSelectionArguments struct {
typeX uint32
}

func (a ardrone3PictureSettingsAutoWhiteBalanceSelection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsAutoWhiteBalanceSelectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var pictureSettingsAutoWhiteBalanceSelection = ardrone3PictureSettingsAutoWhiteBalanceSelection {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdAutoWhiteBalanceSelection,
}

// title : Set image exposure, 
// desc : Set image exposure., 
// support : 0901;090c;090e, 
// result : The exposure is set.\n Then, event [ImageExposure](#1-20-2) is triggered., 
const cmdExpositionSelection cmdDef = 2

type ardrone3PictureSettingsExpositionSelection command

type ardrone3PictureSettingsExpositionSelectionArguments struct {
value float32
}

func (a ardrone3PictureSettingsExpositionSelection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsExpositionSelectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pictureSettingsExpositionSelection = ardrone3PictureSettingsExpositionSelection {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdExpositionSelection,
}

// title : Set image saturation, 
// desc : Set image saturation., 
// support : 0901;090c;090e, 
// result : The saturation is set.\n Then, event [ImageSaturation](#1-20-3) is triggered., 
const cmdSaturationSelection cmdDef = 3

type ardrone3PictureSettingsSaturationSelection command

type ardrone3PictureSettingsSaturationSelectionArguments struct {
value float32
}

func (a ardrone3PictureSettingsSaturationSelection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsSaturationSelectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var pictureSettingsSaturationSelection = ardrone3PictureSettingsSaturationSelection {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdSaturationSelection,
}

// title : Set timelapse mode, 
// desc : Set timelapse mode.\n If timelapse mode is set, instead of taking a video, the drone will take picture regularly.\n Watch out, this command only configure the timelapse mode. Once it is configured, you can start/stop the timelapse with the [RecordVideo](#1-7-3) command., 
// support : 0901;090c;090e, 
// result : The timelapse mode is set (but not started).\n Then, event [TimelapseMode](#1-20-4) is triggered., 
const cmdTimelapseSelection cmdDef = 4

type ardrone3PictureSettingsTimelapseSelection command

type ardrone3PictureSettingsTimelapseSelectionArguments struct {
enabled uint8
interval float32
}

func (a ardrone3PictureSettingsTimelapseSelection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsTimelapseSelectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.enabled)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.interval)
offset += 4

return arg
}

var pictureSettingsTimelapseSelection = ardrone3PictureSettingsTimelapseSelection {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdTimelapseSelection,
}

// title : Set video autorecord mode, 
// desc : Set video autorecord mode.\n If autorecord is set, video record will be automatically started when the drone takes off and stopped slightly after landing., 
// support : 0901;090c;090e, 
// result : The autorecord mode is set.\n Then, event [AutorecordMode](#1-20-5) is triggered., 
const cmdVideoAutorecordSelection cmdDef = 5

type ardrone3PictureSettingsVideoAutorecordSelection command

type ardrone3PictureSettingsVideoAutorecordSelectionArguments struct {
enabled uint8
massstorageid uint8
}

func (a ardrone3PictureSettingsVideoAutorecordSelection) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsVideoAutorecordSelectionArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.enabled)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.massstorageid)
offset++ 

return arg
}

var pictureSettingsVideoAutorecordSelection = ardrone3PictureSettingsVideoAutorecordSelection {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdVideoAutorecordSelection,
}

// title : Set video stabilization mode, 
// desc : Set video stabilization mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video stabilization mode is set.\n Then, event [VideoStabilizationMode](#1-20-6) is triggered., 
const cmdVideoStabilizationMode cmdDef = 6

type ardrone3PictureSettingsVideoStabilizationMode command

type ardrone3PictureSettingsVideoStabilizationModeArguments struct {
mode uint32
}

func (a ardrone3PictureSettingsVideoStabilizationMode) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsVideoStabilizationModeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var pictureSettingsVideoStabilizationMode = ardrone3PictureSettingsVideoStabilizationMode {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdVideoStabilizationMode,
}

// title : Set video recording mode, 
// desc : Set video recording mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video recording mode is set.\n Then, event [VideoRecordingMode](#1-20-7) is triggered., 
const cmdVideoRecordingMode cmdDef = 7

type ardrone3PictureSettingsVideoRecordingMode command

type ardrone3PictureSettingsVideoRecordingModeArguments struct {
mode uint32
}

func (a ardrone3PictureSettingsVideoRecordingMode) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsVideoRecordingModeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var pictureSettingsVideoRecordingMode = ardrone3PictureSettingsVideoRecordingMode {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdVideoRecordingMode,
}

// title : Set video framerate, 
// desc : Set video framerate., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video framerate is set.\n Then, event [VideoFramerate](#1-20-8) is triggered., 
const cmdVideoFramerate cmdDef = 8

type ardrone3PictureSettingsVideoFramerate command

type ardrone3PictureSettingsVideoFramerateArguments struct {
framerate uint32
}

func (a ardrone3PictureSettingsVideoFramerate) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsVideoFramerateArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.framerate)
offset += 4

return arg
}

var pictureSettingsVideoFramerate = ardrone3PictureSettingsVideoFramerate {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdVideoFramerate,
}

// title : Set video resolutions, 
// desc : Set video streaming and recording resolutions., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// result : The video resolutions is set.\n Then, event [VideoResolutions](#1-20-9) is triggered., 
const cmdVideoResolutions cmdDef = 9

type ardrone3PictureSettingsVideoResolutions command

type ardrone3PictureSettingsVideoResolutionsArguments struct {
typeX uint32
}

func (a ardrone3PictureSettingsVideoResolutions) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsVideoResolutionsArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var pictureSettingsVideoResolutions = ardrone3PictureSettingsVideoResolutions {
project: projectardrone3,
class: classPictureSettings,
cmd: cmdVideoResolutions,
}

// Photo settings state from product
const classPictureSettingsState classDef = 20
// title : Picture format, 
// desc : Picture format., 
// support : 0901;090c;090e, 
// triggered : by [SetPictureFormat](#1-19-0)., 
const cmdPictureFormatChanged cmdDef = 0

type ardrone3PictureSettingsStatePictureFormatChanged command

type ardrone3PictureSettingsStatePictureFormatChangedArguments struct {
typeX uint32
}

func (a ardrone3PictureSettingsStatePictureFormatChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStatePictureFormatChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var pictureSettingsStatePictureFormatChanged = ardrone3PictureSettingsStatePictureFormatChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdPictureFormatChanged,
}

// title : White balance mode, 
// desc : White balance mode., 
// support : 0901;090c;090e, 
// triggered : by [SetWhiteBalanceMode](#1-19-1)., 
const cmdAutoWhiteBalanceChanged cmdDef = 1

type ardrone3PictureSettingsStateAutoWhiteBalanceChanged command

type ardrone3PictureSettingsStateAutoWhiteBalanceChangedArguments struct {
typeX uint32
}

func (a ardrone3PictureSettingsStateAutoWhiteBalanceChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateAutoWhiteBalanceChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var pictureSettingsStateAutoWhiteBalanceChanged = ardrone3PictureSettingsStateAutoWhiteBalanceChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdAutoWhiteBalanceChanged,
}

// title : Image exposure, 
// desc : Image exposure., 
// support : 0901;090c;090e, 
// triggered : by [SetImageExposure](#1-19-2)., 
const cmdExpositionChanged cmdDef = 2

type ardrone3PictureSettingsStateExpositionChanged command

type ardrone3PictureSettingsStateExpositionChangedArguments struct {
value float32
min float32
max float32
}

func (a ardrone3PictureSettingsStateExpositionChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateExpositionChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var pictureSettingsStateExpositionChanged = ardrone3PictureSettingsStateExpositionChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdExpositionChanged,
}

// title : Image saturation, 
// desc : Image saturation., 
// support : 0901;090c;090e, 
// triggered : by [SetImageSaturation](#1-19-3)., 
const cmdSaturationChanged cmdDef = 3

type ardrone3PictureSettingsStateSaturationChanged command

type ardrone3PictureSettingsStateSaturationChangedArguments struct {
value float32
min float32
max float32
}

func (a ardrone3PictureSettingsStateSaturationChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateSaturationChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var pictureSettingsStateSaturationChanged = ardrone3PictureSettingsStateSaturationChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdSaturationChanged,
}

// title : Timelapse mode, 
// desc : Timelapse mode., 
// support : 0901;090c;090e, 
// triggered : by [SetTimelapseMode](#1-19-4)., 
const cmdTimelapseChanged cmdDef = 4

type ardrone3PictureSettingsStateTimelapseChanged command

type ardrone3PictureSettingsStateTimelapseChangedArguments struct {
enabled uint8
interval float32
minInterval float32
maxInterval float32
}

func (a ardrone3PictureSettingsStateTimelapseChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateTimelapseChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.enabled)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.interval)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.minInterval)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.maxInterval)
offset += 4

return arg
}

var pictureSettingsStateTimelapseChanged = ardrone3PictureSettingsStateTimelapseChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdTimelapseChanged,
}

// title : Video Autorecord mode, 
// desc : Video Autorecord mode., 
// support : 0901;090c;090e, 
// triggered : by [SetVideoAutorecordMode](#1-19-5)., 
const cmdVideoAutorecordChanged cmdDef = 5

type ardrone3PictureSettingsStateVideoAutorecordChanged command

type ardrone3PictureSettingsStateVideoAutorecordChangedArguments struct {
enabled uint8
massstorageid uint8
}

func (a ardrone3PictureSettingsStateVideoAutorecordChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateVideoAutorecordChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.enabled)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.massstorageid)
offset++ 

return arg
}

var pictureSettingsStateVideoAutorecordChanged = ardrone3PictureSettingsStateVideoAutorecordChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdVideoAutorecordChanged,
}

// title : Video stabilization mode, 
// desc : Video stabilization mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideoStabilizationMode](#1-19-6)., 
const cmdVideoStabilizationModeChanged cmdDef = 6

type ardrone3PictureSettingsStateVideoStabilizationModeChanged command

type ardrone3PictureSettingsStateVideoStabilizationModeChangedArguments struct {
mode uint32
}

func (a ardrone3PictureSettingsStateVideoStabilizationModeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateVideoStabilizationModeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var pictureSettingsStateVideoStabilizationModeChanged = ardrone3PictureSettingsStateVideoStabilizationModeChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdVideoStabilizationModeChanged,
}

// title : Video recording mode, 
// desc : Video recording mode., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideoRecordingMode](#1-19-7)., 
const cmdVideoRecordingModeChanged cmdDef = 7

type ardrone3PictureSettingsStateVideoRecordingModeChanged command

type ardrone3PictureSettingsStateVideoRecordingModeChangedArguments struct {
mode uint32
}

func (a ardrone3PictureSettingsStateVideoRecordingModeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateVideoRecordingModeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var pictureSettingsStateVideoRecordingModeChanged = ardrone3PictureSettingsStateVideoRecordingModeChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdVideoRecordingModeChanged,
}

// title : Video framerate, 
// desc : Video framerate., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideoFramerateMode](#1-19-8)., 
const cmdVideoFramerateChanged cmdDef = 8

type ardrone3PictureSettingsStateVideoFramerateChanged command

type ardrone3PictureSettingsStateVideoFramerateChangedArguments struct {
framerate uint32
}

func (a ardrone3PictureSettingsStateVideoFramerateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateVideoFramerateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.framerate)
offset += 4

return arg
}

var pictureSettingsStateVideoFramerateChanged = ardrone3PictureSettingsStateVideoFramerateChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdVideoFramerateChanged,
}

// title : Video resolutions, 
// desc : Video resolutions.\n This event informs about the recording AND streaming resolutions., 
// support : 0901:3.4.0;090c:3.4.0;090e, 
// triggered : by [SetVideResolutions](#1-19-9)., 
const cmdVideoResolutionsChanged cmdDef = 9

type ardrone3PictureSettingsStateVideoResolutionsChanged command

type ardrone3PictureSettingsStateVideoResolutionsChangedArguments struct {
typeX uint32
}

func (a ardrone3PictureSettingsStateVideoResolutionsChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PictureSettingsStateVideoResolutionsChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var pictureSettingsStateVideoResolutionsChanged = ardrone3PictureSettingsStateVideoResolutionsChanged {
project: projectardrone3,
class: classPictureSettingsState,
cmd: cmdVideoResolutionsChanged,
}

// Control media streaming behavior.
const classMediaStreaming classDef = 21
// title : Enable/disable video streaming, 
// desc : Enable/disable video streaming., 
// support : 0901;090c;090e, 
// result : The video stream is started or stopped.\n Then, event [VideoStreamState](#1-22-0) is triggered., 
const cmdVideoEnable cmdDef = 0

type ardrone3MediaStreamingVideoEnable command

type ardrone3MediaStreamingVideoEnableArguments struct {
enable uint8
}

func (a ardrone3MediaStreamingVideoEnable) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaStreamingVideoEnableArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.enable)
offset++ 

return arg
}

var mediaStreamingVideoEnable = ardrone3MediaStreamingVideoEnable {
project: projectardrone3,
class: classMediaStreaming,
cmd: cmdVideoEnable,
}

// title : Set the stream mode, 
// desc : Set the stream mode., 
// support : 0901;090c;090e, 
// result : The stream mode is set.\n Then, event [VideoStreamMode](#1-22-1) is triggered., 
const cmdVideoStreamMode cmdDef = 1

type ardrone3MediaStreamingVideoStreamMode command

type ardrone3MediaStreamingVideoStreamModeArguments struct {
mode uint32
}

func (a ardrone3MediaStreamingVideoStreamMode) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaStreamingVideoStreamModeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var mediaStreamingVideoStreamMode = ardrone3MediaStreamingVideoStreamMode {
project: projectardrone3,
class: classMediaStreaming,
cmd: cmdVideoStreamMode,
}

// Media streaming status.
const classMediaStreamingState classDef = 22
// title : Video stream state, 
// desc : Video stream state., 
// support : 0901;090c;090e, 
// triggered : by [EnableOrDisableVideoStream](#1-21-0)., 
const cmdVideoEnableChanged cmdDef = 0

type ardrone3MediaStreamingStateVideoEnableChanged command

type ardrone3MediaStreamingStateVideoEnableChangedArguments struct {
enabled uint32
}

func (a ardrone3MediaStreamingStateVideoEnableChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaStreamingStateVideoEnableChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.enabled)
offset += 4

return arg
}

var mediaStreamingStateVideoEnableChanged = ardrone3MediaStreamingStateVideoEnableChanged {
project: projectardrone3,
class: classMediaStreamingState,
cmd: cmdVideoEnableChanged,
}

const cmdVideoStreamModeChanged cmdDef = 1

type ardrone3MediaStreamingStateVideoStreamModeChanged command

type ardrone3MediaStreamingStateVideoStreamModeChangedArguments struct {
mode uint32
}

func (a ardrone3MediaStreamingStateVideoStreamModeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3MediaStreamingStateVideoStreamModeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var mediaStreamingStateVideoStreamModeChanged = ardrone3MediaStreamingStateVideoStreamModeChanged {
project: projectardrone3,
class: classMediaStreamingState,
cmd: cmdVideoStreamModeChanged,
}

// GPS settings
const classGPSSettings classDef = 23
// title : Set home position, 
// desc : Set home position., 
const cmdSetHome cmdDef = 0

type ardrone3GPSSettingsSetHome command

type ardrone3GPSSettingsSetHomeArguments struct {
latitude float64
longitude float64
altitude float64
}

func (a ardrone3GPSSettingsSetHome) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsSetHomeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8

return arg
}

var gPSSettingsSetHome = ardrone3GPSSettingsSetHome {
project: projectardrone3,
class: classGPSSettings,
cmd: cmdSetHome,
}

// title : Reset home position, 
// desc : Reset home position., 
// support : 0901;090c, 
// result : The home position is reset.\n Then, event [HomeLocationReset](#1-24-1) is triggered., 
const cmdResetHome cmdDef = 1

type ardrone3GPSSettingsResetHome command

type ardrone3GPSSettingsResetHomeArguments struct {
}

func (a ardrone3GPSSettingsResetHome) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsResetHomeArguments{}
// No arguments to decode here !!

return arg
}

var gPSSettingsResetHome = ardrone3GPSSettingsResetHome {
project: projectardrone3,
class: classGPSSettings,
cmd: cmdResetHome,
}

// title : Set controller gps location, 
// desc : Set controller gps location.\n The user location might be used in case of return home, according to the home type and the accuracy of the given position. You can get the current home type with the event [HomeType](#1-24-4)., 
// support : 0901;090c;090e, 
// result : The controller position is known by the drone.\n Then, event [HomeLocation](#1-24-2) is triggered., 
const cmdSendControllerGPS cmdDef = 2

type ardrone3GPSSettingsSendControllerGPS command

type ardrone3GPSSettingsSendControllerGPSArguments struct {
latitude float64
longitude float64
altitude float64
horizontalAccuracy float64
verticalAccuracy float64
}

func (a ardrone3GPSSettingsSendControllerGPS) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsSendControllerGPSArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.horizontalAccuracy)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.verticalAccuracy)
offset += 8

return arg
}

var gPSSettingsSendControllerGPS = ardrone3GPSSettingsSendControllerGPS {
project: projectardrone3,
class: classGPSSettings,
cmd: cmdSendControllerGPS,
}

// title : Set the preferred home type, 
// desc : Set the preferred home type.\n Please note that this is only a preference. The actual type chosen is given by the event [HomeType](#1-31-2).\n You can get the currently available types with the event [HomeTypeAvailability](#1-31-1)., 
// support : 0901;090c;090e, 
// result : The user choice is known by the drone.\n Then, event [PreferredHomeType](#1-24-4) is triggered., 
const cmdHomeType cmdDef = 3

type ardrone3GPSSettingsHomeType command

type ardrone3GPSSettingsHomeTypeArguments struct {
typeX uint32
}

func (a ardrone3GPSSettingsHomeType) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsHomeTypeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var gPSSettingsHomeType = ardrone3GPSSettingsHomeType {
project: projectardrone3,
class: classGPSSettings,
cmd: cmdHomeType,
}

// title : Set the return home delay, 
// desc : Set the delay after which the drone will automatically try to return home after a disconnection., 
// support : 0901;090c;090e, 
// result : The delay of the return home is set.\n Then, event [ReturnHomeDelay](#1-24-5) is triggered., 
const cmdReturnHomeDelay cmdDef = 4

type ardrone3GPSSettingsReturnHomeDelay command

type ardrone3GPSSettingsReturnHomeDelayArguments struct {
delay uint16
}

func (a ardrone3GPSSettingsReturnHomeDelay) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsReturnHomeDelayArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.delay)
offset += 2

return arg
}

var gPSSettingsReturnHomeDelay = ardrone3GPSSettingsReturnHomeDelay {
project: projectardrone3,
class: classGPSSettings,
cmd: cmdReturnHomeDelay,
}

// title : Set the return home min altitude, 
// desc : Set the return home minimum altitude. If the drone is below this altitude when starting its return home, it will first reach the minimum altitude. If it is higher than this minimum altitude, it will operate its return home at its actual altitude., 
// support : , 
// result : The minimum altitude for the return home is set.\n Then, event [ReturnHomeMinAltitude](#1-24-7) is triggered., 
const cmdReturnHomeMinAltitude cmdDef = 5

type ardrone3GPSSettingsReturnHomeMinAltitude command

type ardrone3GPSSettingsReturnHomeMinAltitudeArguments struct {
value float32
}

func (a ardrone3GPSSettingsReturnHomeMinAltitude) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsReturnHomeMinAltitudeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4

return arg
}

var gPSSettingsReturnHomeMinAltitude = ardrone3GPSSettingsReturnHomeMinAltitude {
project: projectardrone3,
class: classGPSSettings,
cmd: cmdReturnHomeMinAltitude,
}

// GPS settings state
const classGPSSettingsState classDef = 24
// title : Home location, 
// desc : Home location., 
// support : 0901;090c;090e, 
// triggered : when [HomeType](#1-31-2) changes. Or by [SetHomeLocation](#1-23-2) when [HomeType](#1-31-2) is Pilot. Or regularly after [SetControllerGPS](#140-1) when [HomeType](#1-31-2) is FollowMeTarget. Or at take off [HomeType](#1-31-2) is Takeoff. Or when the first fix occurs and the [HomeType](#1-31-2) is FirstFix., 
const cmdHomeChanged cmdDef = 0

type ardrone3GPSSettingsStateHomeChanged command

type ardrone3GPSSettingsStateHomeChangedArguments struct {
latitude float64
longitude float64
altitude float64
}

func (a ardrone3GPSSettingsStateHomeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateHomeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8

return arg
}

var gPSSettingsStateHomeChanged = ardrone3GPSSettingsStateHomeChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdHomeChanged,
}

// title : Home location has been reset, 
// desc : Home location has been reset., 
// support : 0901;090c, 
// triggered : by [ResetHomeLocation](#1-23-1)., 
const cmdResetHomeChanged cmdDef = 1

type ardrone3GPSSettingsStateResetHomeChanged command

type ardrone3GPSSettingsStateResetHomeChangedArguments struct {
latitude float64
longitude float64
altitude float64
}

func (a ardrone3GPSSettingsStateResetHomeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateResetHomeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.altitude)
offset += 8

return arg
}

var gPSSettingsStateResetHomeChanged = ardrone3GPSSettingsStateResetHomeChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdResetHomeChanged,
}

// title : Gps fix info, 
// desc : Gps fix info., 
// support : 0901;090c;090e, 
// triggered : on change., 
const cmdGPSFixStateChanged cmdDef = 2

type ardrone3GPSSettingsStateGPSFixStateChanged command

type ardrone3GPSSettingsStateGPSFixStateChangedArguments struct {
fixed uint8
}

func (a ardrone3GPSSettingsStateGPSFixStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateGPSFixStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.fixed)
offset++ 

return arg
}

var gPSSettingsStateGPSFixStateChanged = ardrone3GPSSettingsStateGPSFixStateChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdGPSFixStateChanged,
}

// title : Gps update state, 
// desc : Gps update state., 
// support : 0901;090c;090e, 
// triggered : on change., 
const cmdGPSUpdateStateChanged cmdDef = 3

type ardrone3GPSSettingsStateGPSUpdateStateChanged command

type ardrone3GPSSettingsStateGPSUpdateStateChangedArguments struct {
state uint32
}

func (a ardrone3GPSSettingsStateGPSUpdateStateChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateGPSUpdateStateChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4

return arg
}

var gPSSettingsStateGPSUpdateStateChanged = ardrone3GPSSettingsStateGPSUpdateStateChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdGPSUpdateStateChanged,
}

// title : Preferred home type, 
// desc : User preference for the home type.\n See [HomeType](#1-31-2) to get the drone actual home type., 
// support : 0901;090c;090e, 
// triggered : by [SetPreferredHomeType](#1-23-3)., 
const cmdHomeTypeChanged cmdDef = 4

type ardrone3GPSSettingsStateHomeTypeChanged command

type ardrone3GPSSettingsStateHomeTypeChangedArguments struct {
typeX uint32
}

func (a ardrone3GPSSettingsStateHomeTypeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateHomeTypeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var gPSSettingsStateHomeTypeChanged = ardrone3GPSSettingsStateHomeTypeChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdHomeTypeChanged,
}

// title : Return home delay, 
// desc : Return home trigger delay. This delay represents the time after which the return home is automatically triggered after a disconnection., 
// support : 0901;090c;090e, 
// triggered : by [SetReturnHomeDelay](#1-23-4)., 
const cmdReturnHomeDelayChanged cmdDef = 5

type ardrone3GPSSettingsStateReturnHomeDelayChanged command

type ardrone3GPSSettingsStateReturnHomeDelayChangedArguments struct {
delay uint16
}

func (a ardrone3GPSSettingsStateReturnHomeDelayChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateReturnHomeDelayChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+2]), binary.LittleEndian, &arg.delay)
offset += 2

return arg
}

var gPSSettingsStateReturnHomeDelayChanged = ardrone3GPSSettingsStateReturnHomeDelayChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdReturnHomeDelayChanged,
}

// title : Geofence center, 
// desc : Geofence center location. This location represents the center of the geofence zone. This is updated at a maximum frequency of 1 Hz., 
// triggered : when [HomeChanged](#1-24-0) and when [GpsLocationChanged](#1-4-9) before takeoff., 
const cmdGeofenceCenterChanged cmdDef = 6

type ardrone3GPSSettingsStateGeofenceCenterChanged command

type ardrone3GPSSettingsStateGeofenceCenterChangedArguments struct {
latitude float64
longitude float64
}

func (a ardrone3GPSSettingsStateGeofenceCenterChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateGeofenceCenterChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.latitude)
offset += 8
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.longitude)
offset += 8

return arg
}

var gPSSettingsStateGeofenceCenterChanged = ardrone3GPSSettingsStateGeofenceCenterChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdGeofenceCenterChanged,
}

// title : Return home min altitude, 
// desc : Minumum altitude for return home changed., 
// triggered : by [SetReturnHomeMinAltitude](#1-23-5)., 
const cmdReturnHomeMinAltitudeChanged cmdDef = 7

type ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged command

type ardrone3GPSSettingsStateReturnHomeMinAltitudeChangedArguments struct {
value float32
min float32
max float32
}

func (a ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSSettingsStateReturnHomeMinAltitudeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.value)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.min)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.max)
offset += 4

return arg
}

var gPSSettingsStateReturnHomeMinAltitudeChanged = ardrone3GPSSettingsStateReturnHomeMinAltitudeChanged {
project: projectardrone3,
class: classGPSSettingsState,
cmd: cmdReturnHomeMinAltitudeChanged,
}

// Camera state
const classCameraState classDef = 25
// title : Camera orientation, 
// desc : Camera orientation., 
// support : 0901;090c;090e, 
// triggered : by [SetCameraOrientation](#1-1-0)., 
const cmdOrientationDUPLICATE cmdDef = 0

type ardrone3CameraStateOrientation command

type ardrone3CameraStateOrientationArguments struct {
tilt int8
pan int8
}

func (a ardrone3CameraStateOrientation) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraStateOrientationArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.tilt)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.pan)
offset++ 

return arg
}

var cameraStateOrientation = ardrone3CameraStateOrientation {
project: projectardrone3,
class: classCameraState,
cmd: cmdOrientation,
}

// title : Orientation of the camera center, 
// desc : Orientation of the center of the camera.\n This is the value to send when you want to center the camera., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const cmddefaultCameraOrientation cmdDef = 1

type ardrone3CameraStatedefaultCameraOrientation command

type ardrone3CameraStatedefaultCameraOrientationArguments struct {
tilt int8
pan int8
}

func (a ardrone3CameraStatedefaultCameraOrientation) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraStatedefaultCameraOrientationArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.tilt)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.pan)
offset++ 

return arg
}

var cameraStatedefaultCameraOrientation = ardrone3CameraStatedefaultCameraOrientation {
project: projectardrone3,
class: classCameraState,
cmd: cmddefaultCameraOrientation,
}

// title : Camera orientation, 
// desc : Camera orientation with float arguments., 
// support : 0901;090c;090e, 
// triggered : by [SetCameraOrientationV2](#1-1-1), 
const cmdOrientationV2DUPLICATE cmdDef = 2

type ardrone3CameraStateOrientationV2 command

type ardrone3CameraStateOrientationV2Arguments struct {
tilt float32
pan float32
}

func (a ardrone3CameraStateOrientationV2) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraStateOrientationV2Arguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.tilt)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.pan)
offset += 4

return arg
}

var cameraStateOrientationV2 = ardrone3CameraStateOrientationV2 {
project: projectardrone3,
class: classCameraState,
cmd: cmdOrientationV2,
}

// title : Orientation of the camera center, 
// desc : Orientation of the center of the camera.\n This is the value to send when you want to center the camera., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const cmddefaultCameraOrientationV2 cmdDef = 3

type ardrone3CameraStatedefaultCameraOrientationV2 command

type ardrone3CameraStatedefaultCameraOrientationV2Arguments struct {
tilt float32
pan float32
}

func (a ardrone3CameraStatedefaultCameraOrientationV2) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraStatedefaultCameraOrientationV2Arguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.tilt)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.pan)
offset += 4

return arg
}

var cameraStatedefaultCameraOrientationV2 = ardrone3CameraStatedefaultCameraOrientationV2 {
project: projectardrone3,
class: classCameraState,
cmd: cmddefaultCameraOrientationV2,
}

// title : Camera velocity range, 
// desc : Camera Orientation velocity limits., 
// support : 0901;090c;090e, 
// triggered : at connection., 
const cmdVelocityRange cmdDef = 4

type ardrone3CameraStateVelocityRange command

type ardrone3CameraStateVelocityRangeArguments struct {
maxtilt float32
maxpan float32
}

func (a ardrone3CameraStateVelocityRange) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3CameraStateVelocityRangeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.maxtilt)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.maxpan)
offset += 4

return arg
}

var cameraStateVelocityRange = ardrone3CameraStateVelocityRange {
project: projectardrone3,
class: classCameraState,
cmd: cmdVelocityRange,
}

// Anti-flickering related commands
const classAntiflickering classDef = 29
// title : Set the electric frequency, 
// desc : Set the electric frequency of the surrounding lights.\n This is used to avoid the video flickering in auto mode. You can get the current antiflickering mode with the event [AntiflickeringModeChanged](#1-30-1)., 
// support : 0901;090c, 
// result : The electric frequency is set.\n Then, event [ElectricFrequency](#1-30-0) is triggered., 
const cmdelectricFrequency cmdDef = 0

type ardrone3AntiflickeringelectricFrequency command

type ardrone3AntiflickeringelectricFrequencyArguments struct {
frequency uint32
}

func (a ardrone3AntiflickeringelectricFrequency) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3AntiflickeringelectricFrequencyArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.frequency)
offset += 4

return arg
}

var antiflickeringelectricFrequency = ardrone3AntiflickeringelectricFrequency {
project: projectardrone3,
class: classAntiflickering,
cmd: cmdelectricFrequency,
}

// title : Set the antiflickering mode, 
// desc : Set the antiflickering mode.\n If auto, the drone will detect when flickers appears on the video and trigger the antiflickering.\n In this case, this electric frequency it will use will be the one specified in the event [ElectricFrequency](#1-29-0).\n Forcing the antiflickering (FixedFiftyHertz or FixedFiftyHertz) can reduce luminosity of the video., 
// support : 0901;090c, 
// result : The antiflickering mode is set.\n Then, event [AntiflickeringMode](#1-30-1) is triggered., 
const cmdsetMode cmdDef = 1

type ardrone3AntiflickeringsetMode command

type ardrone3AntiflickeringsetModeArguments struct {
mode uint32
}

func (a ardrone3AntiflickeringsetMode) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3AntiflickeringsetModeArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var antiflickeringsetMode = ardrone3AntiflickeringsetMode {
project: projectardrone3,
class: classAntiflickering,
cmd: cmdsetMode,
}

// Anti-flickering related states
const classAntiflickeringState classDef = 30
// title : Electric frequency, 
// desc : Electric frequency.\n This piece of information is used for the antiflickering when the [AntiflickeringMode](#1-30-1) is set to *auto*., 
// support : 0901;090c, 
// triggered : by [SetElectricFrequency](#1-29-0)., 
const cmdelectricFrequencyChanged cmdDef = 0

type ardrone3AntiflickeringStateelectricFrequencyChanged command

type ardrone3AntiflickeringStateelectricFrequencyChangedArguments struct {
frequency uint32
}

func (a ardrone3AntiflickeringStateelectricFrequencyChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3AntiflickeringStateelectricFrequencyChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.frequency)
offset += 4

return arg
}

var antiflickeringStateelectricFrequencyChanged = ardrone3AntiflickeringStateelectricFrequencyChanged {
project: projectardrone3,
class: classAntiflickeringState,
cmd: cmdelectricFrequencyChanged,
}

// title : Antiflickering mode, 
// desc : Antiflickering mode., 
// support : 0901;090c, 
// triggered : by [SetAntiflickeringMode](#1-29-1)., 
const cmdmodeChanged cmdDef = 1

type ardrone3AntiflickeringStatemodeChanged command

type ardrone3AntiflickeringStatemodeChangedArguments struct {
mode uint32
}

func (a ardrone3AntiflickeringStatemodeChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3AntiflickeringStatemodeChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.mode)
offset += 4

return arg
}

var antiflickeringStatemodeChanged = ardrone3AntiflickeringStatemodeChanged {
project: projectardrone3,
class: classAntiflickeringState,
cmd: cmdmodeChanged,
}

// GPS related States
const classGPSState classDef = 31
// title : Number of GPS satellites, 
// desc : Number of GPS satellites., 
// support : 0901;090c;090e, 
// triggered : on change., 
const cmdNumberOfSatelliteChanged cmdDef = 0

type ardrone3GPSStateNumberOfSatelliteChanged command

type ardrone3GPSStateNumberOfSatelliteChangedArguments struct {
numberOfSatellite uint8
}

func (a ardrone3GPSStateNumberOfSatelliteChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSStateNumberOfSatelliteChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.numberOfSatellite)
offset++ 

return arg
}

var gPSStateNumberOfSatelliteChanged = ardrone3GPSStateNumberOfSatelliteChanged {
project: projectardrone3,
class: classGPSState,
cmd: cmdNumberOfSatelliteChanged,
}

// title : Home type availability, 
// desc : Home type availability., 
// support : 0901;090c;090e, 
// triggered : when the availability of, at least, one type changes.\n This might be due to controller position availability, gps fix before take off or other reason., 
const cmdHomeTypeAvailabilityChanged cmdDef = 1

type ardrone3GPSStateHomeTypeAvailabilityChanged command

type ardrone3GPSStateHomeTypeAvailabilityChangedArguments struct {
typeX uint32
available uint8
}

func (a ardrone3GPSStateHomeTypeAvailabilityChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSStateHomeTypeAvailabilityChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.available)
offset++ 

return arg
}

var gPSStateHomeTypeAvailabilityChanged = ardrone3GPSStateHomeTypeAvailabilityChanged {
project: projectardrone3,
class: classGPSState,
cmd: cmdHomeTypeAvailabilityChanged,
}

// title : Home type, 
// desc : Home type.\n This choice is made by the drone, according to the [PreferredHomeType](#1-24-4) and the [HomeTypeAvailability](#1-31-1). The drone will choose the type matching with the user preference only if this type is available. If not, it will chose a type in this order:\n FOLLOWEE ; TAKEOFF ; PILOT ; FIRST_FIX, 
// support : 0901;090c;090e, 
// triggered : when the return home type chosen by the drone changes.\n This might be produced by a user preference triggered by [SetPreferedHomeType](#1-23-3) or by a change in the [HomeTypesAvailabilityChanged](#1-31-1)., 
const cmdHomeTypeChosenChanged cmdDef = 2

type ardrone3GPSStateHomeTypeChosenChanged command

type ardrone3GPSStateHomeTypeChosenChangedArguments struct {
typeX uint32
}

func (a ardrone3GPSStateHomeTypeChosenChanged) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3GPSStateHomeTypeChosenChangedArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.typeX)
offset += 4

return arg
}

var gPSStateHomeTypeChosenChanged = ardrone3GPSStateHomeTypeChosenChanged {
project: projectardrone3,
class: classGPSState,
cmd: cmdHomeTypeChosenChanged,
}

// Pro features enabled on the Bebop
const classPROState classDef = 32
// title : Pro features, 
// desc : Pro features., 
const cmdFeatures cmdDef = 0

type ardrone3PROStateFeatures command

type ardrone3PROStateFeaturesArguments struct {
features uint64
}

func (a ardrone3PROStateFeatures) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3PROStateFeaturesArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+8]), binary.LittleEndian, &arg.features)
offset += 8

return arg
}

var pROStateFeatures = ardrone3PROStateFeatures {
project: projectardrone3,
class: classPROState,
cmd: cmdFeatures,
}

// Information about the connected accessories
const classAccessoryState classDef = 33
// title : List of connected accessories, 
// desc : List of all connected accessories. This event presents the list of all connected accessories. To actually use the component, use the component dedicated feature., 
// support : 090e:1.5.0, 
// triggered : at connection or when an accessory is connected., 
const cmdConnectedAccessories cmdDef = 0

type ardrone3AccessoryStateConnectedAccessories command

type ardrone3AccessoryStateConnectedAccessoriesArguments struct {
id uint8
accessorytype uint32
uid string
swVersion string
listflags uint8
}

func (a ardrone3AccessoryStateConnectedAccessories) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
var stringEnd int
var err error
arg := ardrone3AccessoryStateConnectedAccessoriesArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.id)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.accessorytype)
offset += 4

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.uid = string(b[offset:offset+stringEnd])
offset += stringEnd

				stringEnd, err = getLengthOfStringData(b[offset:])
				if err != nil {
					log.Println("error: ", err)
				}
arg.swVersion = string(b[offset:offset+stringEnd])
offset += stringEnd
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.listflags)
offset++ 

return arg
}

var accessoryStateConnectedAccessories = ardrone3AccessoryStateConnectedAccessories {
project: projectardrone3,
class: classAccessoryState,
cmd: cmdConnectedAccessories,
}

// title : Connected accessories battery, 
// desc : Connected accessories battery., 
// support : none, 
const cmdBattery cmdDef = 1

type ardrone3AccessoryStateBattery command

type ardrone3AccessoryStateBatteryArguments struct {
id uint8
batteryLevel uint8
listflags uint8
}

func (a ardrone3AccessoryStateBattery) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3AccessoryStateBatteryArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.id)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.batteryLevel)
offset++ 
binary.Read(bytes.NewReader(b[offset:offset+1]), binary.LittleEndian, &arg.listflags)
offset++ 

return arg
}

var accessoryStateBattery = ardrone3AccessoryStateBattery {
project: projectardrone3,
class: classAccessoryState,
cmd: cmdBattery,
}

// Sounds related commands
const classSound classDef = 35
// title : Start alert sound, 
// desc : Start the alert sound. The alert sound can only be started when the drone is not flying., 
// support : none, 
// result : The drone makes a sound and send back [AlertSoundState](#1-36-0) with state playing., 
const cmdStartAlertSound cmdDef = 0

type ardrone3SoundStartAlertSound command

type ardrone3SoundStartAlertSoundArguments struct {
}

func (a ardrone3SoundStartAlertSound) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SoundStartAlertSoundArguments{}
// No arguments to decode here !!

return arg
}

var soundStartAlertSound = ardrone3SoundStartAlertSound {
project: projectardrone3,
class: classSound,
cmd: cmdStartAlertSound,
}

// title : Stop alert sound, 
// desc : Stop the alert sound., 
// support : none, 
// result : The drone stops its alert sound and send back [AlertSoundState](#1-36-0) with state stopped., 
const cmdStopAlertSound cmdDef = 1

type ardrone3SoundStopAlertSound command

type ardrone3SoundStopAlertSoundArguments struct {
}

func (a ardrone3SoundStopAlertSound) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SoundStopAlertSoundArguments{}
// No arguments to decode here !!

return arg
}

var soundStopAlertSound = ardrone3SoundStopAlertSound {
project: projectardrone3,
class: classSound,
cmd: cmdStopAlertSound,
}

// Sounds related events
const classSoundState classDef = 36
// title : Alert sound state, 
// desc : Alert sound state., 
// support : none, 
// triggered : by [StartAlertSound](#1-35-0) or [StopAlertSound](#1-35-1) or when the drone starts or stops to play an alert sound by itself., 
const cmdAlertSound cmdDef = 0

type ardrone3SoundStateAlertSound command

type ardrone3SoundStateAlertSoundArguments struct {
state uint32
}

func (a ardrone3SoundStateAlertSound) decode(b []byte) interface{} {
//TODO: .............
fmt.Printf(".....we are now decoding the payload %v, which is of type %T\n", a, a)
fmt.Printf("%+v\n", a)
arg := ardrone3SoundStateAlertSoundArguments{}
var offset = 0
binary.Read(bytes.NewReader(b[offset:offset+4]), binary.LittleEndian, &arg.state)
offset += 4

return arg
}

var soundStateAlertSound = ardrone3SoundStateAlertSound {
project: projectardrone3,
class: classSoundState,
cmd: cmdAlertSound,
}

type decoder interface {
decode([]byte) interface{}
}

var commandMap = map[command]decoder {
command(pilotingTakeOff) : pilotingTakeOff,
command(pilotingPCMD) : pilotingPCMD,
command(pilotingLanding) : pilotingLanding,
command(pilotingEmergency) : pilotingEmergency,
command(pilotingNavigateHome) : pilotingNavigateHome,
command(pilotingAutoTakeOffMode) : pilotingAutoTakeOffMode,
command(pilotingmoveBy) : pilotingmoveBy,
command(pilotingUserTakeOff) : pilotingUserTakeOff,
command(pilotingCircle) : pilotingCircle,
command(pilotingmoveTo) : pilotingmoveTo,
command(pilotingCancelMoveTo) : pilotingCancelMoveTo,
command(pilotingStartPilotedPOI) : pilotingStartPilotedPOI,
command(pilotingStopPilotedPOI) : pilotingStopPilotedPOI,
command(pilotingCancelMoveBy) : pilotingCancelMoveBy,
command(animationsFlip) : animationsFlip,
command(cameraOrientation) : cameraOrientation,
command(cameraOrientationV2) : cameraOrientationV2,
command(cameraVelocity) : cameraVelocity,
command(mediaRecordPicture) : mediaRecordPicture,
command(mediaRecordVideo) : mediaRecordVideo,
command(mediaRecordPictureV2) : mediaRecordPictureV2,
command(mediaRecordVideoV2) : mediaRecordVideoV2,
command(mediaRecordStatePictureStateChanged) : mediaRecordStatePictureStateChanged,
command(mediaRecordStateVideoStateChanged) : mediaRecordStateVideoStateChanged,
command(mediaRecordStatePictureStateChangedV2) : mediaRecordStatePictureStateChangedV2,
command(mediaRecordStateVideoStateChangedV2) : mediaRecordStateVideoStateChangedV2,
command(mediaRecordStateVideoResolutionState) : mediaRecordStateVideoResolutionState,
command(mediaRecordEventPictureEventChanged) : mediaRecordEventPictureEventChanged,
command(mediaRecordEventVideoEventChanged) : mediaRecordEventVideoEventChanged,
command(pilotingStateFlyingStateChanged) : pilotingStateFlyingStateChanged,
command(pilotingStateAlertStateChanged) : pilotingStateAlertStateChanged,
command(pilotingStateNavigateHomeStateChanged) : pilotingStateNavigateHomeStateChanged,
command(pilotingStatePositionChanged) : pilotingStatePositionChanged,
command(pilotingStateSpeedChanged) : pilotingStateSpeedChanged,
command(pilotingStateAttitudeChanged) : pilotingStateAttitudeChanged,
command(pilotingStateAutoTakeOffModeChanged) : pilotingStateAutoTakeOffModeChanged,
command(pilotingStateAltitudeChanged) : pilotingStateAltitudeChanged,
command(pilotingStateGpsLocationChanged) : pilotingStateGpsLocationChanged,
command(pilotingStateLandingStateChanged) : pilotingStateLandingStateChanged,
command(pilotingStateAirSpeedChanged) : pilotingStateAirSpeedChanged,
command(pilotingStatemoveToChanged) : pilotingStatemoveToChanged,
command(pilotingStateMotionState) : pilotingStateMotionState,
command(pilotingStatePilotedPOI) : pilotingStatePilotedPOI,
command(pilotingStateReturnHomeBatteryCapacity) : pilotingStateReturnHomeBatteryCapacity,
command(pilotingStatemoveByChanged) : pilotingStatemoveByChanged,
command(pilotingStateHoveringWarning) : pilotingStateHoveringWarning,
command(pilotingStateForcedLandingAutoTrigger) : pilotingStateForcedLandingAutoTrigger,
command(pilotingStateWindStateChanged) : pilotingStateWindStateChanged,
command(pilotingEventmoveByEnd) : pilotingEventmoveByEnd,
command(networkWifiScan) : networkWifiScan,
command(networkWifiAuthChannel) : networkWifiAuthChannel,
command(networkStateWifiScanListChanged) : networkStateWifiScanListChanged,
command(networkStateAllWifiScanChanged) : networkStateAllWifiScanChanged,
command(networkStateWifiAuthChannelListChanged) : networkStateWifiAuthChannelListChanged,
command(networkStateAllWifiAuthChannelChanged) : networkStateAllWifiAuthChannelChanged,
command(pilotingSettingsMaxAltitude) : pilotingSettingsMaxAltitude,
command(pilotingSettingsMaxTilt) : pilotingSettingsMaxTilt,
command(pilotingSettingsAbsolutControl) : pilotingSettingsAbsolutControl,
command(pilotingSettingsMaxDistance) : pilotingSettingsMaxDistance,
command(pilotingSettingsNoFlyOverMaxDistance) : pilotingSettingsNoFlyOverMaxDistance,
command(pilotingSettingssetAutonomousFlightMaxHorizontalSpeed) : pilotingSettingssetAutonomousFlightMaxHorizontalSpeed,
command(pilotingSettingssetAutonomousFlightMaxVerticalSpeed) : pilotingSettingssetAutonomousFlightMaxVerticalSpeed,
command(pilotingSettingssetAutonomousFlightMaxHorizontalAcceleration) : pilotingSettingssetAutonomousFlightMaxHorizontalAcceleration,
command(pilotingSettingssetAutonomousFlightMaxVerticalAcceleration) : pilotingSettingssetAutonomousFlightMaxVerticalAcceleration,
command(pilotingSettingssetAutonomousFlightMaxRotationSpeed) : pilotingSettingssetAutonomousFlightMaxRotationSpeed,
command(pilotingSettingsBankedTurn) : pilotingSettingsBankedTurn,
command(pilotingSettingsMinAltitude) : pilotingSettingsMinAltitude,
command(pilotingSettingsCirclingDirection) : pilotingSettingsCirclingDirection,
command(pilotingSettingsCirclingRadius) : pilotingSettingsCirclingRadius,
command(pilotingSettingsCirclingAltitude) : pilotingSettingsCirclingAltitude,
command(pilotingSettingsPitchMode) : pilotingSettingsPitchMode,
command(pilotingSettingsSetMotionDetectionMode) : pilotingSettingsSetMotionDetectionMode,
command(pilotingSettingsStateMaxAltitudeChanged) : pilotingSettingsStateMaxAltitudeChanged,
command(pilotingSettingsStateMaxTiltChanged) : pilotingSettingsStateMaxTiltChanged,
command(pilotingSettingsStateAbsolutControlChanged) : pilotingSettingsStateAbsolutControlChanged,
command(pilotingSettingsStateMaxDistanceChanged) : pilotingSettingsStateMaxDistanceChanged,
command(pilotingSettingsStateNoFlyOverMaxDistanceChanged) : pilotingSettingsStateNoFlyOverMaxDistanceChanged,
command(pilotingSettingsStateAutonomousFlightMaxHorizontalSpeed) : pilotingSettingsStateAutonomousFlightMaxHorizontalSpeed,
command(pilotingSettingsStateAutonomousFlightMaxVerticalSpeed) : pilotingSettingsStateAutonomousFlightMaxVerticalSpeed,
command(pilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration) : pilotingSettingsStateAutonomousFlightMaxHorizontalAcceleration,
command(pilotingSettingsStateAutonomousFlightMaxVerticalAcceleration) : pilotingSettingsStateAutonomousFlightMaxVerticalAcceleration,
command(pilotingSettingsStateAutonomousFlightMaxRotationSpeed) : pilotingSettingsStateAutonomousFlightMaxRotationSpeed,
command(pilotingSettingsStateBankedTurnChanged) : pilotingSettingsStateBankedTurnChanged,
command(pilotingSettingsStateMinAltitudeChanged) : pilotingSettingsStateMinAltitudeChanged,
command(pilotingSettingsStateCirclingDirectionChanged) : pilotingSettingsStateCirclingDirectionChanged,
command(pilotingSettingsStateCirclingRadiusChanged) : pilotingSettingsStateCirclingRadiusChanged,
command(pilotingSettingsStateCirclingAltitudeChanged) : pilotingSettingsStateCirclingAltitudeChanged,
command(pilotingSettingsStatePitchModeChanged) : pilotingSettingsStatePitchModeChanged,
command(pilotingSettingsStateMotionDetection) : pilotingSettingsStateMotionDetection,
command(speedSettingsMaxVerticalSpeed) : speedSettingsMaxVerticalSpeed,
command(speedSettingsMaxRotationSpeed) : speedSettingsMaxRotationSpeed,
command(speedSettingsHullProtection) : speedSettingsHullProtection,
command(speedSettingsOutdoor) : speedSettingsOutdoor,
command(speedSettingsMaxPitchRollRotationSpeed) : speedSettingsMaxPitchRollRotationSpeed,
command(speedSettingsStateMaxVerticalSpeedChanged) : speedSettingsStateMaxVerticalSpeedChanged,
command(speedSettingsStateMaxRotationSpeedChanged) : speedSettingsStateMaxRotationSpeedChanged,
command(speedSettingsStateHullProtectionChanged) : speedSettingsStateHullProtectionChanged,
command(speedSettingsStateOutdoorChanged) : speedSettingsStateOutdoorChanged,
command(speedSettingsStateMaxPitchRollRotationSpeedChanged) : speedSettingsStateMaxPitchRollRotationSpeedChanged,
command(networkSettingsWifiSelection) : networkSettingsWifiSelection,
command(networkSettingswifiSecurity) : networkSettingswifiSecurity,
command(networkSettingsStateWifiSelectionChanged) : networkSettingsStateWifiSelectionChanged,
command(networkSettingsStatewifiSecurityChanged) : networkSettingsStatewifiSecurityChanged,
command(networkSettingsStatewifiSecurity) : networkSettingsStatewifiSecurity,
command(settingsStateProductMotorVersionListChanged) : settingsStateProductMotorVersionListChanged,
command(settingsStateProductGPSVersionChanged) : settingsStateProductGPSVersionChanged,
command(settingsStateMotorErrorStateChanged) : settingsStateMotorErrorStateChanged,
command(settingsStateMotorSoftwareVersionChanged) : settingsStateMotorSoftwareVersionChanged,
command(settingsStateMotorFlightsStatusChanged) : settingsStateMotorFlightsStatusChanged,
command(settingsStateMotorErrorLastErrorChanged) : settingsStateMotorErrorLastErrorChanged,
command(settingsStateP7ID) : settingsStateP7ID,
command(settingsStateCPUID) : settingsStateCPUID,
command(pictureSettingsPictureFormatSelection) : pictureSettingsPictureFormatSelection,
command(pictureSettingsAutoWhiteBalanceSelection) : pictureSettingsAutoWhiteBalanceSelection,
command(pictureSettingsExpositionSelection) : pictureSettingsExpositionSelection,
command(pictureSettingsSaturationSelection) : pictureSettingsSaturationSelection,
command(pictureSettingsTimelapseSelection) : pictureSettingsTimelapseSelection,
command(pictureSettingsVideoAutorecordSelection) : pictureSettingsVideoAutorecordSelection,
command(pictureSettingsVideoStabilizationMode) : pictureSettingsVideoStabilizationMode,
command(pictureSettingsVideoRecordingMode) : pictureSettingsVideoRecordingMode,
command(pictureSettingsVideoFramerate) : pictureSettingsVideoFramerate,
command(pictureSettingsVideoResolutions) : pictureSettingsVideoResolutions,
command(pictureSettingsStatePictureFormatChanged) : pictureSettingsStatePictureFormatChanged,
command(pictureSettingsStateAutoWhiteBalanceChanged) : pictureSettingsStateAutoWhiteBalanceChanged,
command(pictureSettingsStateExpositionChanged) : pictureSettingsStateExpositionChanged,
command(pictureSettingsStateSaturationChanged) : pictureSettingsStateSaturationChanged,
command(pictureSettingsStateTimelapseChanged) : pictureSettingsStateTimelapseChanged,
command(pictureSettingsStateVideoAutorecordChanged) : pictureSettingsStateVideoAutorecordChanged,
command(pictureSettingsStateVideoStabilizationModeChanged) : pictureSettingsStateVideoStabilizationModeChanged,
command(pictureSettingsStateVideoRecordingModeChanged) : pictureSettingsStateVideoRecordingModeChanged,
command(pictureSettingsStateVideoFramerateChanged) : pictureSettingsStateVideoFramerateChanged,
command(pictureSettingsStateVideoResolutionsChanged) : pictureSettingsStateVideoResolutionsChanged,
command(mediaStreamingVideoEnable) : mediaStreamingVideoEnable,
command(mediaStreamingVideoStreamMode) : mediaStreamingVideoStreamMode,
command(mediaStreamingStateVideoEnableChanged) : mediaStreamingStateVideoEnableChanged,
command(mediaStreamingStateVideoStreamModeChanged) : mediaStreamingStateVideoStreamModeChanged,
command(gPSSettingsSetHome) : gPSSettingsSetHome,
command(gPSSettingsResetHome) : gPSSettingsResetHome,
command(gPSSettingsSendControllerGPS) : gPSSettingsSendControllerGPS,
command(gPSSettingsHomeType) : gPSSettingsHomeType,
command(gPSSettingsReturnHomeDelay) : gPSSettingsReturnHomeDelay,
command(gPSSettingsReturnHomeMinAltitude) : gPSSettingsReturnHomeMinAltitude,
command(gPSSettingsStateHomeChanged) : gPSSettingsStateHomeChanged,
command(gPSSettingsStateResetHomeChanged) : gPSSettingsStateResetHomeChanged,
command(gPSSettingsStateGPSFixStateChanged) : gPSSettingsStateGPSFixStateChanged,
command(gPSSettingsStateGPSUpdateStateChanged) : gPSSettingsStateGPSUpdateStateChanged,
command(gPSSettingsStateHomeTypeChanged) : gPSSettingsStateHomeTypeChanged,
command(gPSSettingsStateReturnHomeDelayChanged) : gPSSettingsStateReturnHomeDelayChanged,
command(gPSSettingsStateGeofenceCenterChanged) : gPSSettingsStateGeofenceCenterChanged,
command(gPSSettingsStateReturnHomeMinAltitudeChanged) : gPSSettingsStateReturnHomeMinAltitudeChanged,
command(cameraStateOrientation) : cameraStateOrientation,
command(cameraStatedefaultCameraOrientation) : cameraStatedefaultCameraOrientation,
command(cameraStateOrientationV2) : cameraStateOrientationV2,
command(cameraStatedefaultCameraOrientationV2) : cameraStatedefaultCameraOrientationV2,
command(cameraStateVelocityRange) : cameraStateVelocityRange,
command(antiflickeringelectricFrequency) : antiflickeringelectricFrequency,
command(antiflickeringsetMode) : antiflickeringsetMode,
command(antiflickeringStateelectricFrequencyChanged) : antiflickeringStateelectricFrequencyChanged,
command(antiflickeringStatemodeChanged) : antiflickeringStatemodeChanged,
command(gPSStateNumberOfSatelliteChanged) : gPSStateNumberOfSatelliteChanged,
command(gPSStateHomeTypeAvailabilityChanged) : gPSStateHomeTypeAvailabilityChanged,
command(gPSStateHomeTypeChosenChanged) : gPSStateHomeTypeChosenChanged,
command(pROStateFeatures) : pROStateFeatures,
command(accessoryStateConnectedAccessories) : accessoryStateConnectedAccessories,
command(accessoryStateBattery) : accessoryStateBattery,
command(soundStartAlertSound) : soundStartAlertSound,
command(soundStopAlertSound) : soundStopAlertSound,
command(soundStateAlertSound) : soundStateAlertSound,
}


	// lenStringData takes a []byte which is the data for the arguments, and returns
	// the position of the 0 terminator for the string.
	// The []byte given as input will start looking from the beginning of the slice,
	// so the input slice should be sliced to start from the offset of the string.
	func lenStringData(b []byte) (int, error) {
		// Figure out the length of the string
		for i := 0; i < cap(b); i++ {
			//fmt.Printf("%+v, of type %T\n", b[i], b[i])

			//fmt.Println("i = ", i)
			if b[i] == 0 {
				//fmt.Println("lengthString = ", i)

				// add 1 to jump to the 0
				return i + 1, nil
			}

		}

		err := fmt.Errorf("no string bytes found, returning 0")
		return 0, err
	}
	

	func getLengthOfStringData(b []byte) (int, error) {
		// Figure out the length of the string
		for i := 0; i < cap(b); i++ {
			//fmt.Printf("%+v, of type %T\n", b[i], b[i])
	
			//fmt.Println("i = ", i)
			if b[i] == 0 {
				//fmt.Println("lengthString = ", i)
	
				// add 1 to jump to the 0
				return i + 1, nil
			}
	
		}
	
		err := fmt.Errorf("no string bytes found, returning 0")
		return 0, err
	}
	
