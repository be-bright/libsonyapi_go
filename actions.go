package main

type Action struct {
	SetShootMode                string
	GetShootMode                string
	GetSupportedShootMode       string
	GetAvailableShootMode       string
	ActTakePicture              string
	AwaitTakePicture            string
	StartContShooting           string
	StopContShooting            string
	StartMovieRec               string
	StopMovieRec                string
	StartAudioRec               string
	StopAudioRec                string
	StartIntervalStillRec       string
	StopIntervalStillRec        string
	StartLoopRec                string
	StopLoopRec                 string
	StartLiveview               string
	StopLiveview                string
	StartLiveviewWithSize       string
	GetLiveviewSize             string
	GetSupportedLiveviewSize    string
	GetAvailableLiveviewSize    string
	SetLiveviewFrameInfo        string
	GetLiveviewFrameInfo        string
	ActZoom                     string
	SetZoomSetting              string
	GetZoomSetting              string
	GetSupportedZoomSetting     string
	GetAvailableZoomSetting     string
	ActHalfPressShutter         string
	CancelHalfPressShutter      string
	SetTouchAFPosition          string
	GetTouchAFPosition          string
	CancelTouchAFPosition       string
	ActTrackingFocus            string
	CancelTrackingFocus         string
	SetTrackingFocus            string
	GetTrackingFocus            string
	GetSupportedTrackingFocus   string
	GetAvailableTrackingFocus   string
	SetContShootingMode         string
	GetContShootingMode         string
	GetSupportedContShootingMode string
	GetAvailableContShootingMode string
	setContShootingSpeed 		string
    getContShootingSpeed 		string
    getSupportedContShootingSpeed string
    getAvailableContShootingSpeed string
    setSelfTimer 				string
    getSelfTimer 				string
    getSupportedSelfTimer 		string
    getAvailableSelfTimer 		string
    setExposureMode 			string
    getExposureMode 			string
    getSupportedExposureMode 	string
    getAvailableExposureMode 	string
    setFocusMode 				string
    getFocusMode 				string
    getSupportedFocusMode 		string
    getAvailableFocusMode 		string
    setExposureCompensation 	string
    getExposureCompensation 	string
    getSupportedExposureCompensation string
    getAvailableExposureCompensation string
    setFNumber 					string
    getFNumber 					string
    getSupportedFNumber 		string
    getAvailableFNumber 		string
    setShutterSpeed 			string
    getShutterSpeed 			string
    getSupportedShutterSpeed 	string
    getAvailableShutterSpeed 	string
    setIsoSpeedRate 			string
    getIsoSpeedRate 			string
    getSupportedIsoSpeedRate 	string
    getAvailableIsoSpeedRate 	string
    setWhiteBalance 			string
    getWhiteBalance 			string
    getSupportedWhiteBalance 	string
    getAvailableWhiteBalance 	string
    actWhiteBalanceOnePushCustom string
    setProgramShift 			string
    getSupportedProgramShift 	string
    setFlashMode 				string
    getFlashMode 				string
    getSupportedFlashMode 		string
    getAvailableFlashMode 		string
    setStillSize 				string
    getStillSize 				string
    getSupportedStillSize 		string
    getAvailableStillSize 		string
    setStillQuality 			string
    getStillQuality 			string
    getSupportedStillQuality 	string
    getAvailableStillQuality 	string
    setPostviewImageSize 		string
    getPostviewImageSize 		string
    getSupportedPostviewImageSize string
    getAvailablePostviewImageSize string
    setMovieFileFormat 			string
    getMovieFileFormat 			string
    getSupportedMovieFileFormat string
    getAvailableMovieFileFormat string
    setMovieQuality 			string
    getMovieQuality 			string
    getSupportedMovieQuality 	string
    getAvailableMovieQuality 	string
    setSteadyMode 				string
    getSteadyMode 				string
    getSupportedSteadyMode 		string
    getAvailableSteadyMode 		string
    setViewAngle 				string
    getViewAngle 				string
    getSupportedViewAngle 		string
    getAvailableViewAngle 		string
    setSceneSelection 			string
    getSceneSelection 			string
    getSupportedSceneSelection 	string
    getAvailableSceneSelection 	string
    setColorSetting 			string
    getColorSetting 			string
    getSupportedColorSetting 	string
    getAvailableColorSetting 	string
    setIntervalTime 			string
    getIntervalTime 			string
    getSupportedIntervalTime 	string
    getAvailableIntervalTime 	string
    setLoopRecTime 				string
    getLoopRecTime 				string
    getSupportedLoopRecTime 	string
    getAvailableLoopRecTime 	string
    setWindNoiseReduction 		string
    getWindNoiseReduction 		string
    getSupportedWindNoiseReduction string
    getAvailableWindNoiseReduction string
    setAudioRecording 			string
    getAudioRecording 			string
    getSupportedAudioRecording 	string
    getAvailableAudioRecording 	string
    setFlipSetting 				string
    getFlipSetting 				string
    getSupportedFlipSetting 	string
    getAvailableFlipSetting 	string
    setTvColorSystem 			string
    getTvColorSystem 			string
    getSupportedTvColorSystem 	string
    getAvailableTvColorSystem 	string
    startRecMode 				string
    stopRecMo 					string
}

func Actions() *Action {
	return &Actions{
		SetShootMode:                "setShootMode",
		GetShootMode:                "getShootMode",
		GetSupportedShootMode:       "getSupportedShootMode",
		GetAvailableShootMode:       "getAvailableShootMode",
		ActTakePicture:              "actTakePicture",
		AwaitTakePicture:            "awaitTakePicture",
		StartContShooting:           "startContShooting",
		StopContShooting:            "stopContShooting",
		StartMovieRec:               "startMovieRec",
		StopMovieRec:                "stopMovieRec",
		StartAudioRec:               "startAudioRec",
		StopAudioRec:                "stopAudioRec",
		StartIntervalStillRec:       "startIntervalStillRec",
		StopIntervalStillRec:        "stopIntervalStillRec",
		StartLoopRec:                "startLoopRec",
		StopLoopRec:                 "stopLoopRec",
		StartLiveview:               "startLiveview",
		StopLiveview:                "stopLiveview",
		StartLiveviewWithSize:       "startLiveviewWithSize",
		GetLiveviewSize:             "getLiveviewSize",
		GetSupportedLiveviewSize:    "getSupportedLiveviewSize",
		GetAvailableLiveviewSize:    "getAvailableLiveviewSize",
		SetLiveviewFrameInfo:        "setLiveviewFrameInfo",
		GetLiveviewFrameInfo:        "getLiveviewFrameInfo",
		ActZoom:                     "actZoom",
		SetZoomSetting:              "setZoomSetting",
		GetZoomSetting:              "getZoomSetting",
		GetSupportedZoomSetting:     "getSupportedZoomSetting",
		GetAvailableZoomSetting:     "getAvailableZoomSetting",
		ActHalfPressShutter:         "actHalfPressShutter",
		CancelHalfPressShutter:      "cancelHalfPressShutter",
		SetTouchAFPosition:          "setTouchAFPosition",
		GetTouchAFPosition:          "getTouchAFPosition",
		CancelTouchAFPosition:       "cancelTouchAFPosition",
		ActTrackingFocus:            "actTrackingFocus",
		CancelTrackingFocus:         "cancelTrackingFocus",
		SetTrackingFocus:            "setTrackingFocus",
		GetTrackingFocus:            "getTrackingFocus",
		GetSupportedTrackingFocus:   "getSupportedTrackingFocus",
		GetAvailableTrackingFocus:   "getAvailableTrackingFocus",
		SetContShootingMode:         "setContShootingMode",
		GetContShootingMode:         "getContShootingMode",
		GetSupportedContShootingMode: "getSupportedContShootingMode",
		GetAvailableContShootingMode: "getAvailableContShootingMode",
		setContShootingSpeed: 		  "setContShootingSpeed",
		getContShootingSpeed: 		  "getContShootingSpeed",
		getSupportedContShootingSpeed:"getSupportedContShootingSpeed",
		getAvailableContShootingSpeed:"getAvailableContShootingSpeed",
		setSelfTimer:                "setSelfTimer",
		getSelfTimer:                "getSelfTimer",
		getSupportedSelfTimer:       "getSupportedSelfTimer",
		getAvailableSelfTimer:       "getAvailableSelfTimer",
		setExposureMode: "setExposureMode",
		getExposureMode: "getExposureMode",
		getSupportedExposureMode: "getSupportedExposureMode",
		getAvailableExposureMode: "getAvailableExposureMode",
		setFocusMode: "setFocusMode",
		getFocusMode: "getFocusMode",
		getSupportedFocusMode: "getSupportedFocusMode",
		getAvailableFocusMode: "getAvailableFocusMode",
		setExposureCompensation: "setExposureCompensation",
		getExposureCompensation: "getExposureCompensation",
		getSupportedExposureCompensation: "getSupportedExposureCompensation",
		getAvailableExposureCompensation: "getAvailableExposureCompensation",
		setFNumber: "setFNumber",
		getFNumber: "getFNumber",
		getSupportedFNumber: "getSupportedFNumber",
		getAvailableFNumber: "getAvailableFNumber",
		setShutterSpeed: "setShutterSpeed",
		getShutterSpeed: "getShutterSpeed",
		getSupportedShutterSpeed: "getSupportedShutterSpeed",
		getAvailableShutterSpeed: "getAvailableShutterSpeed",
		setIsoSpeedRate: "setIsoSpeedRate",
		getIsoSpeedRate: "getIsoSpeedRate",
		getSupportedIsoSpeedRate: "getSupportedIsoSpeedRate",
		getAvailableIsoSpeedRate: "getAvailableIsoSpeedRate",
		setWhiteBalance: "setWhiteBalance",
		getWhiteBalance: "getWhiteBalance",
		getSupportedWhiteBalance: "getSupportedWhiteBalance",
		getAvailableWhiteBalance: "getAvailableWhiteBalance",
		actWhiteBalanceOnePushCustom: "actWhiteBalanceOnePushCustom",
		setProgramShift: "setProgramShift",
		getSupportedProgramShift: "getSupportedProgramShift",
		setFlashMode: "setFlashMode",
		getFlashMode: "getFlashMode",
		getSupportedFlashMode: "getSupportedFlashMode",
		getAvailableFlashMode: "getAvailableFlashMode",
		setStillSize: "setStillSize",
		getStillSize: "getStillSize",
		getSupportedStillSize: "getSupportedStillSize",
		getAvailableStillSize: "getAvailableStillSize",
		setStillQuality: "setStillQuality",
		getStillQuality: "getStillQuality",
		getSupportedStillQuality: "getSupportedStillQuality",
		getAvailableStillQuality: "getAvailableStillQuality",
		setPostviewImageSize: "setPostviewImageSize",
		getPostviewImageSize: "getPostviewImageSize",
		getSupportedPostviewImageSize: "getSupportedPostviewImageSize",
		getAvailablePostviewImageSize: "getAvailablePostviewImageSize",
		setMovieFileFormat: "setMovieFileFormat",
		getMovieFileFormat: "getMovieFileFormat",
		getSupportedMovieFileFormat: "getSupportedMovieFileFormat",
		getAvailableMovieFileFormat: "getAvailableMovieFileFormat",
		setMovieQuality: "setMovieQuality",
		getMovieQuality: "getMovieQuality",
		getSupportedMovieQuality: "getSupportedMovieQuality",
		getAvailableMovieQuality: "getAvailableMovieQuality",
		setSteadyMode: "setSteadyMode",
		getSteadyMode: "getSteadyMode",
		getSupportedSteadyMode: "getSupportedSteadyMode",
		getAvailableSteadyMode: "getAvailableSteadyMode",
		setViewAngle: "setViewAngle",
		getViewAngle: "getViewAngle",
		getSupportedViewAngle: "getSupportedViewAngle",
		getAvailableViewAngle: "getAvailableViewAngle",
		setSceneSelection: "setSceneSelection",
		getSceneSelection: "getSceneSelection",
		getSupportedSceneSelection: "getSupportedSceneSelection",
		getAvailableSceneSelection: "getAvailableSceneSelection",
		setColorSetting: "setColorSetting",
		getColorSetting: "getColorSetting",
		getSupportedColorSetting: "getSupportedColorSetting",
		getAvailableColorSetting: "getAvailableColorSetting",
		setIntervalTime: "setIntervalTime",
		getIntervalTime: "getIntervalTime",
		getSupportedIntervalTime: "getSupportedIntervalTime",
		getAvailableIntervalTime: "getAvailableIntervalTime",
		setLoopRecTime: "setLoopRecTime",
		getLoopRecTime: "getLoopRecTime",
		getSupportedLoopRecTime: "getSupportedLoopRecTime",
		getAvailableLoopRecTime: "getAvailableLoopRecTime",
		setWindNoiseReduction: "setWindNoiseReduction",
		getWindNoiseReduction: "getWindNoiseReduction",
		getSupportedWindNoiseReduction: "getSupportedWindNoiseReduction",
		getAvailableWindNoiseReduction: "getAvailableWindNoiseReduction",
		setAudioRecording: "setAudioRecording",
		getAudioRecording: "getAudioRecording",
		getSupportedAudioRecording: "getSupportedAudioRecording",
		getAvailableAudioRecording: "getAvailableAudioRecording",
		setFlipSetting: "setFlipSetting",
		getFlipSetting: "getFlipSetting",
		getSupportedFlipSetting: "getSupportedFlipSetting",
		getAvailableFlipSetting: "getAvailableFlipSetting",
		setTvColorSystem: "setTvColorSystem",
		getTvColorSystem: "getTvColorSystem",
		getSupportedTvColorSystem: "getSupportedTvColorSystem",
		getAvailableTvColorSystem: "getAvailableTvColorSystem",
		startRecMode: "startRecMode",
		stopRecMo: "stopRecMo",
	}
}

func main() {
	actions := Actions()
	// usage example:
	// fmt.Println(actions.SetShootMode)
	// ...
}
