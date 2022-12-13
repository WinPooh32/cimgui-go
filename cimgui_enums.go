package cimgui

// original name: ImDrawFlags_
type DrawFlags int

const (
	DrawFlags_None                    = 0
	DrawFlags_Closed                  = 1
	DrawFlags_RoundCornersTopLeft     = 16
	DrawFlags_RoundCornersTopRight    = 32
	DrawFlags_RoundCornersBottomLeft  = 64
	DrawFlags_RoundCornersBottomRight = 128
	DrawFlags_RoundCornersNone        = 256
	DrawFlags_RoundCornersTop         = 48
	DrawFlags_RoundCornersBottom      = 192
	DrawFlags_RoundCornersLeft        = 80
	DrawFlags_RoundCornersRight       = 160
	DrawFlags_RoundCornersAll         = 240
	DrawFlags_RoundCornersDefault_    = 240
	DrawFlags_RoundCornersMask_       = 496
)

// original name: ImDrawListFlags_
type DrawListFlags int

const (
	DrawListFlags_None                   = 0
	DrawListFlags_AntiAliasedLines       = 1
	DrawListFlags_AntiAliasedLinesUseTex = 2
	DrawListFlags_AntiAliasedFill        = 4
	DrawListFlags_AllowVtxOffset         = 8
)

// original name: ImFontAtlasFlags_
type FontAtlasFlags int

const (
	FontAtlasFlags_None               = 0
	FontAtlasFlags_NoPowerOfTwoHeight = 1
	FontAtlasFlags_NoMouseCursors     = 2
	FontAtlasFlags_NoBakedLines       = 4
)

// original name: ImGuiActivateFlags_
type ActivateFlags int

const (
	ActivateFlags_None               = 0
	ActivateFlags_PreferInput        = 1
	ActivateFlags_PreferTweak        = 2
	ActivateFlags_TryToPreserveState = 4
)

// original name: ImGuiAxis
type Axis int

const (
	Axis_None = -1
	Axis_X    = 0
	Axis_Y    = 1
)

// original name: ImGuiBackendFlags_
type BackendFlags int

const (
	BackendFlags_None                    = 0
	BackendFlags_HasGamepad              = 1
	BackendFlags_HasMouseCursors         = 2
	BackendFlags_HasSetMousePos          = 4
	BackendFlags_RendererHasVtxOffset    = 8
	BackendFlags_PlatformHasViewports    = 1024
	BackendFlags_HasMouseHoveredViewport = 2048
	BackendFlags_RendererHasViewports    = 4096
)

// original name: ImGuiButtonFlagsPrivate_
type ButtonFlagsPrivate int

const (
	ButtonFlags_PressedOnClick                = 16
	ButtonFlags_PressedOnClickRelease         = 32
	ButtonFlags_PressedOnClickReleaseAnywhere = 64
	ButtonFlags_PressedOnRelease              = 128
	ButtonFlags_PressedOnDoubleClick          = 256
	ButtonFlags_PressedOnDragDropHold         = 512
	ButtonFlags_Repeat                        = 1024
	ButtonFlags_FlattenChildren               = 2048
	ButtonFlags_AllowItemOverlap              = 4096
	ButtonFlags_DontClosePopups               = 8192
	ButtonFlags_AlignTextBaseLine             = 32768
	ButtonFlags_NoKeyModifiers                = 65536
	ButtonFlags_NoHoldingActiveId             = 131072
	ButtonFlags_NoNavFocus                    = 262144
	ButtonFlags_NoHoveredOnFocus              = 524288
	ButtonFlags_PressedOnMask_                = 1008
	ButtonFlags_PressedOnDefault_             = 32
)

// original name: ImGuiButtonFlags_
type ButtonFlags int

const (
	ButtonFlags_None                = 0
	ButtonFlags_MouseButtonLeft     = 1
	ButtonFlags_MouseButtonRight    = 2
	ButtonFlags_MouseButtonMiddle   = 4
	ButtonFlags_MouseButtonMask_    = 7
	ButtonFlags_MouseButtonDefault_ = 1
)

// original name: ImGuiCol_
type Col int

const (
	Col_Text                  = 0
	Col_TextDisabled          = 1
	Col_WindowBg              = 2
	Col_ChildBg               = 3
	Col_PopupBg               = 4
	Col_Border                = 5
	Col_BorderShadow          = 6
	Col_FrameBg               = 7
	Col_FrameBgHovered        = 8
	Col_FrameBgActive         = 9
	Col_TitleBg               = 10
	Col_TitleBgActive         = 11
	Col_TitleBgCollapsed      = 12
	Col_MenuBarBg             = 13
	Col_ScrollbarBg           = 14
	Col_ScrollbarGrab         = 15
	Col_ScrollbarGrabHovered  = 16
	Col_ScrollbarGrabActive   = 17
	Col_CheckMark             = 18
	Col_SliderGrab            = 19
	Col_SliderGrabActive      = 20
	Col_Button                = 21
	Col_ButtonHovered         = 22
	Col_ButtonActive          = 23
	Col_Header                = 24
	Col_HeaderHovered         = 25
	Col_HeaderActive          = 26
	Col_Separator             = 27
	Col_SeparatorHovered      = 28
	Col_SeparatorActive       = 29
	Col_ResizeGrip            = 30
	Col_ResizeGripHovered     = 31
	Col_ResizeGripActive      = 32
	Col_Tab                   = 33
	Col_TabHovered            = 34
	Col_TabActive             = 35
	Col_TabUnfocused          = 36
	Col_TabUnfocusedActive    = 37
	Col_DockingPreview        = 38
	Col_DockingEmptyBg        = 39
	Col_PlotLines             = 40
	Col_PlotLinesHovered      = 41
	Col_PlotHistogram         = 42
	Col_PlotHistogramHovered  = 43
	Col_TableHeaderBg         = 44
	Col_TableBorderStrong     = 45
	Col_TableBorderLight      = 46
	Col_TableRowBg            = 47
	Col_TableRowBgAlt         = 48
	Col_TextSelectedBg        = 49
	Col_DragDropTarget        = 50
	Col_NavHighlight          = 51
	Col_NavWindowingHighlight = 52
	Col_NavWindowingDimBg     = 53
	Col_ModalWindowDimBg      = 54
	Col_COUNT                 = 55
)

// original name: ImGuiColorEditFlags_
type ColorEditFlags int

const (
	ColorEditFlags_None             = 0
	ColorEditFlags_NoAlpha          = 2
	ColorEditFlags_NoPicker         = 4
	ColorEditFlags_NoOptions        = 8
	ColorEditFlags_NoSmallPreview   = 16
	ColorEditFlags_NoInputs         = 32
	ColorEditFlags_NoTooltip        = 64
	ColorEditFlags_NoLabel          = 128
	ColorEditFlags_NoSidePreview    = 256
	ColorEditFlags_NoDragDrop       = 512
	ColorEditFlags_NoBorder         = 1024
	ColorEditFlags_AlphaBar         = 65536
	ColorEditFlags_AlphaPreview     = 131072
	ColorEditFlags_AlphaPreviewHalf = 262144
	ColorEditFlags_HDR              = 524288
	ColorEditFlags_DisplayRGB       = 1048576
	ColorEditFlags_DisplayHSV       = 2097152
	ColorEditFlags_DisplayHex       = 4194304
	ColorEditFlags_Uint8            = 8388608
	ColorEditFlags_Float            = 16777216
	ColorEditFlags_PickerHueBar     = 33554432
	ColorEditFlags_PickerHueWheel   = 67108864
	ColorEditFlags_InputRGB         = 134217728
	ColorEditFlags_InputHSV         = 268435456
	ColorEditFlags_DefaultOptions_  = 177209344
	ColorEditFlags_DisplayMask_     = 7340032
	ColorEditFlags_DataTypeMask_    = 25165824
	ColorEditFlags_PickerMask_      = 100663296
	ColorEditFlags_InputMask_       = 402653184
)

// original name: ImGuiComboFlagsPrivate_
type ComboFlagsPrivate int

const (
	ComboFlags_CustomPreview = 1048576
)

// original name: ImGuiComboFlags_
type ComboFlags int

const (
	ComboFlags_None           = 0
	ComboFlags_PopupAlignLeft = 1
	ComboFlags_HeightSmall    = 2
	ComboFlags_HeightRegular  = 4
	ComboFlags_HeightLarge    = 8
	ComboFlags_HeightLargest  = 16
	ComboFlags_NoArrowButton  = 32
	ComboFlags_NoPreview      = 64
	ComboFlags_HeightMask_    = 30
)

// original name: ImGuiCond_
type Cond int

const (
	Cond_None         = 0
	Cond_Always       = 1
	Cond_Once         = 2
	Cond_FirstUseEver = 4
	Cond_Appearing    = 8
)

// original name: ImGuiConfigFlags_
type ConfigFlags int

const (
	ConfigFlags_None                    = 0
	ConfigFlags_NavEnableKeyboard       = 1
	ConfigFlags_NavEnableGamepad        = 2
	ConfigFlags_NavEnableSetMousePos    = 4
	ConfigFlags_NavNoCaptureKeyboard    = 8
	ConfigFlags_NoMouse                 = 16
	ConfigFlags_NoMouseCursorChange     = 32
	ConfigFlags_DockingEnable           = 64
	ConfigFlags_ViewportsEnable         = 1024
	ConfigFlags_DpiEnableScaleViewports = 16384
	ConfigFlags_DpiEnableScaleFonts     = 32768
	ConfigFlags_IsSRGB                  = 1048576
	ConfigFlags_IsTouchScreen           = 2097152
)

// original name: ImGuiContextHookType
type ContextHookType int

const (
	ContextHookType_NewFramePre     = 0
	ContextHookType_NewFramePost    = 1
	ContextHookType_EndFramePre     = 2
	ContextHookType_EndFramePost    = 3
	ContextHookType_RenderPre       = 4
	ContextHookType_RenderPost      = 5
	ContextHookType_Shutdown        = 6
	ContextHookType_PendingRemoval_ = 7
)

// original name: ImGuiDataAuthority_
type DataAuthority int

const (
	DataAuthority_Auto     = 0
	DataAuthority_DockNode = 1
	DataAuthority_Window   = 2
)

// original name: ImGuiDataTypePrivate_
type DataTypePrivate int

const (
	DataType_String  = 11
	DataType_Pointer = 12
	DataType_ID      = 13
)

// original name: ImGuiDataType_
type DataType int

const (
	DataType_S8     = 0
	DataType_U8     = 1
	DataType_S16    = 2
	DataType_U16    = 3
	DataType_S32    = 4
	DataType_U32    = 5
	DataType_S64    = 6
	DataType_U64    = 7
	DataType_Float  = 8
	DataType_Double = 9
	DataType_COUNT  = 10
)

// original name: ImGuiDebugLogFlags_
type DebugLogFlags int

const (
	DebugLogFlags_None          = 0
	DebugLogFlags_EventActiveId = 1
	DebugLogFlags_EventFocus    = 2
	DebugLogFlags_EventPopup    = 4
	DebugLogFlags_EventNav      = 8
	DebugLogFlags_EventClipper  = 16
	DebugLogFlags_EventIO       = 32
	DebugLogFlags_EventDocking  = 64
	DebugLogFlags_EventViewport = 128
	DebugLogFlags_EventMask_    = 255
	DebugLogFlags_OutputToTTY   = 1024
)

// original name: ImGuiDir_
type Dir int

const (
	Dir_None  = -1
	Dir_Left  = 0
	Dir_Right = 1
	Dir_Up    = 2
	Dir_Down  = 3
	Dir_COUNT = 4
)

// original name: ImGuiDockNodeFlagsPrivate_
type DockNodeFlagsPrivate int

const (
	DockNodeFlags_DockSpace               = 1024
	DockNodeFlags_CentralNode             = 2048
	DockNodeFlags_NoTabBar                = 4096
	DockNodeFlags_HiddenTabBar            = 8192
	DockNodeFlags_NoWindowMenuButton      = 16384
	DockNodeFlags_NoCloseButton           = 32768
	DockNodeFlags_NoDocking               = 65536
	DockNodeFlags_NoDockingSplitMe        = 131072
	DockNodeFlags_NoDockingSplitOther     = 262144
	DockNodeFlags_NoDockingOverMe         = 524288
	DockNodeFlags_NoDockingOverOther      = 1048576
	DockNodeFlags_NoDockingOverEmpty      = 2097152
	DockNodeFlags_NoResizeX               = 4194304
	DockNodeFlags_NoResizeY               = 8388608
	DockNodeFlags_SharedFlagsInheritMask_ = -1
	DockNodeFlags_NoResizeFlagsMask_      = 12582944
	DockNodeFlags_LocalFlagsMask_         = 12713072
	DockNodeFlags_LocalFlagsTransferMask_ = 12712048
	DockNodeFlags_SavedFlagsMask_         = 12712992
)

// original name: ImGuiDockNodeFlags_
type DockNodeFlags int

const (
	DockNodeFlags_None                   = 0
	DockNodeFlags_KeepAliveOnly          = 1
	DockNodeFlags_NoDockingInCentralNode = 4
	DockNodeFlags_PassthruCentralNode    = 8
	DockNodeFlags_NoSplit                = 16
	DockNodeFlags_NoResize               = 32
	DockNodeFlags_AutoHideTabBar         = 64
)

// original name: ImGuiDockNodeState
type DockNodeState int

const (
	DockNodeState_Unknown                                   = 0
	DockNodeState_HostWindowHiddenBecauseSingleWindow       = 1
	DockNodeState_HostWindowHiddenBecauseWindowsAreResizing = 2
	DockNodeState_HostWindowVisible                         = 3
)

// original name: ImGuiDragDropFlags_
type DragDropFlags int

const (
	DragDropFlags_None                     = 0
	DragDropFlags_SourceNoPreviewTooltip   = 1
	DragDropFlags_SourceNoDisableHover     = 2
	DragDropFlags_SourceNoHoldToOpenOthers = 4
	DragDropFlags_SourceAllowNullID        = 8
	DragDropFlags_SourceExtern             = 16
	DragDropFlags_SourceAutoExpirePayload  = 32
	DragDropFlags_AcceptBeforeDelivery     = 1024
	DragDropFlags_AcceptNoDrawDefaultRect  = 2048
	DragDropFlags_AcceptNoPreviewTooltip   = 4096
	DragDropFlags_AcceptPeekOnly           = 3072
)

// original name: ImGuiFocusedFlags_
type FocusedFlags int

const (
	FocusedFlags_None                = 0
	FocusedFlags_ChildWindows        = 1
	FocusedFlags_RootWindow          = 2
	FocusedFlags_AnyWindow           = 4
	FocusedFlags_NoPopupHierarchy    = 8
	FocusedFlags_DockHierarchy       = 16
	FocusedFlags_RootAndChildWindows = 3
)

// original name: ImGuiHoveredFlags_
type HoveredFlags int

const (
	HoveredFlags_None                         = 0
	HoveredFlags_ChildWindows                 = 1
	HoveredFlags_RootWindow                   = 2
	HoveredFlags_AnyWindow                    = 4
	HoveredFlags_NoPopupHierarchy             = 8
	HoveredFlags_DockHierarchy                = 16
	HoveredFlags_AllowWhenBlockedByPopup      = 32
	HoveredFlags_AllowWhenBlockedByActiveItem = 128
	HoveredFlags_AllowWhenOverlapped          = 256
	HoveredFlags_AllowWhenDisabled            = 512
	HoveredFlags_NoNavOverride                = 1024
	HoveredFlags_RectOnly                     = 416
	HoveredFlags_RootAndChildWindows          = 3
)

// original name: ImGuiInputEventType
type InputEventType int

const (
	InputEventType_None          = 0
	InputEventType_MousePos      = 1
	InputEventType_MouseWheel    = 2
	InputEventType_MouseButton   = 3
	InputEventType_MouseViewport = 4
	InputEventType_Key           = 5
	InputEventType_Text          = 6
	InputEventType_Focus         = 7
	InputEventType_COUNT         = 8
)

// original name: ImGuiInputFlags_
type InputFlags int

const (
	InputFlags_None               = 0
	InputFlags_Repeat             = 1
	InputFlags_RepeatRateDefault  = 2
	InputFlags_RepeatRateNavMove  = 4
	InputFlags_RepeatRateNavTweak = 8
	InputFlags_RepeatRateMask_    = 14
)

// original name: ImGuiInputSource
type InputSource int

const (
	InputSource_None      = 0
	InputSource_Mouse     = 1
	InputSource_Keyboard  = 2
	InputSource_Gamepad   = 3
	InputSource_Clipboard = 4
	InputSource_Nav       = 5
	InputSource_COUNT     = 6
)

// original name: ImGuiInputTextFlagsPrivate_
type InputTextFlagsPrivate int

const (
	InputTextFlags_Multiline    = 67108864
	InputTextFlags_NoMarkEdited = 134217728
	InputTextFlags_MergedItem   = 268435456
)

// original name: ImGuiInputTextFlags_
type InputTextFlags int

const (
	InputTextFlags_None                = 0
	InputTextFlags_CharsDecimal        = 1
	InputTextFlags_CharsHexadecimal    = 2
	InputTextFlags_CharsUppercase      = 4
	InputTextFlags_CharsNoBlank        = 8
	InputTextFlags_AutoSelectAll       = 16
	InputTextFlags_EnterReturnsTrue    = 32
	InputTextFlags_CallbackCompletion  = 64
	InputTextFlags_CallbackHistory     = 128
	InputTextFlags_CallbackAlways      = 256
	InputTextFlags_CallbackCharFilter  = 512
	InputTextFlags_AllowTabInput       = 1024
	InputTextFlags_CtrlEnterForNewLine = 2048
	InputTextFlags_NoHorizontalScroll  = 4096
	InputTextFlags_AlwaysOverwrite     = 8192
	InputTextFlags_ReadOnly            = 16384
	InputTextFlags_Password            = 32768
	InputTextFlags_NoUndoRedo          = 65536
	InputTextFlags_CharsScientific     = 131072
	InputTextFlags_CallbackResize      = 262144
	InputTextFlags_CallbackEdit        = 524288
)

// original name: ImGuiItemFlags_
type ItemFlags int

const (
	ItemFlags_None                     = 0
	ItemFlags_NoTabStop                = 1
	ItemFlags_ButtonRepeat             = 2
	ItemFlags_Disabled                 = 4
	ItemFlags_NoNav                    = 8
	ItemFlags_NoNavDefaultFocus        = 16
	ItemFlags_SelectableDontClosePopup = 32
	ItemFlags_MixedValue               = 64
	ItemFlags_ReadOnly                 = 128
	ItemFlags_Inputable                = 256
)

// original name: ImGuiItemStatusFlags_
type ItemStatusFlags int

const (
	ItemStatusFlags_None             = 0
	ItemStatusFlags_HoveredRect      = 1
	ItemStatusFlags_HasDisplayRect   = 2
	ItemStatusFlags_Edited           = 4
	ItemStatusFlags_ToggledSelection = 8
	ItemStatusFlags_ToggledOpen      = 16
	ItemStatusFlags_HasDeactivated   = 32
	ItemStatusFlags_Deactivated      = 64
	ItemStatusFlags_HoveredWindow    = 128
	ItemStatusFlags_FocusedByTabbing = 256
)

// original name: ImGuiKeyPrivate_
type KeyPrivate int

const (
	Key_LegacyNativeKey_BEGIN = 0
	Key_LegacyNativeKey_END   = 512
	Key_Keyboard_BEGIN        = 512
	Key_Keyboard_END          = 617
	Key_Gamepad_BEGIN         = 617
	Key_Gamepad_END           = 641
	Key_Aliases_BEGIN         = 645
	Key_Aliases_END           = 652
	Key_NavKeyboardTweakSlow  = 641
	Key_NavKeyboardTweakFast  = 642
	Key_NavGamepadTweakSlow   = 627
	Key_NavGamepadTweakFast   = 628
	Key_NavGamepadActivate    = 622
	Key_NavGamepadCancel      = 620
	Key_NavGamepadMenu        = 619
	Key_NavGamepadInput       = 621
)

// original name: ImGuiKey_
type Key int

const (
	Key_None               = 0
	Key_Tab                = 512
	Key_LeftArrow          = 513
	Key_RightArrow         = 514
	Key_UpArrow            = 515
	Key_DownArrow          = 516
	Key_PageUp             = 517
	Key_PageDown           = 518
	Key_Home               = 519
	Key_End                = 520
	Key_Insert             = 521
	Key_Delete             = 522
	Key_Backspace          = 523
	Key_Space              = 524
	Key_Enter              = 525
	Key_Escape             = 526
	Key_LeftCtrl           = 527
	Key_LeftShift          = 528
	Key_LeftAlt            = 529
	Key_LeftSuper          = 530
	Key_RightCtrl          = 531
	Key_RightShift         = 532
	Key_RightAlt           = 533
	Key_RightSuper         = 534
	Key_Menu               = 535
	Key_0                  = 536
	Key_1                  = 537
	Key_2                  = 538
	Key_3                  = 539
	Key_4                  = 540
	Key_5                  = 541
	Key_6                  = 542
	Key_7                  = 543
	Key_8                  = 544
	Key_9                  = 545
	Key_A                  = 546
	Key_B                  = 547
	Key_C                  = 548
	Key_D                  = 549
	Key_E                  = 550
	Key_F                  = 551
	Key_G                  = 552
	Key_H                  = 553
	Key_I                  = 554
	Key_J                  = 555
	Key_K                  = 556
	Key_L                  = 557
	Key_M                  = 558
	Key_N                  = 559
	Key_O                  = 560
	Key_P                  = 561
	Key_Q                  = 562
	Key_R                  = 563
	Key_S                  = 564
	Key_T                  = 565
	Key_U                  = 566
	Key_V                  = 567
	Key_W                  = 568
	Key_X                  = 569
	Key_Y                  = 570
	Key_Z                  = 571
	Key_F1                 = 572
	Key_F2                 = 573
	Key_F3                 = 574
	Key_F4                 = 575
	Key_F5                 = 576
	Key_F6                 = 577
	Key_F7                 = 578
	Key_F8                 = 579
	Key_F9                 = 580
	Key_F10                = 581
	Key_F11                = 582
	Key_F12                = 583
	Key_Apostrophe         = 584
	Key_Comma              = 585
	Key_Minus              = 586
	Key_Period             = 587
	Key_Slash              = 588
	Key_Semicolon          = 589
	Key_Equal              = 590
	Key_LeftBracket        = 591
	Key_Backslash          = 592
	Key_RightBracket       = 593
	Key_GraveAccent        = 594
	Key_CapsLock           = 595
	Key_ScrollLock         = 596
	Key_NumLock            = 597
	Key_PrintScreen        = 598
	Key_Pause              = 599
	Key_Keypad0            = 600
	Key_Keypad1            = 601
	Key_Keypad2            = 602
	Key_Keypad3            = 603
	Key_Keypad4            = 604
	Key_Keypad5            = 605
	Key_Keypad6            = 606
	Key_Keypad7            = 607
	Key_Keypad8            = 608
	Key_Keypad9            = 609
	Key_KeypadDecimal      = 610
	Key_KeypadDivide       = 611
	Key_KeypadMultiply     = 612
	Key_KeypadSubtract     = 613
	Key_KeypadAdd          = 614
	Key_KeypadEnter        = 615
	Key_KeypadEqual        = 616
	Key_GamepadStart       = 617
	Key_GamepadBack        = 618
	Key_GamepadFaceLeft    = 619
	Key_GamepadFaceRight   = 620
	Key_GamepadFaceUp      = 621
	Key_GamepadFaceDown    = 622
	Key_GamepadDpadLeft    = 623
	Key_GamepadDpadRight   = 624
	Key_GamepadDpadUp      = 625
	Key_GamepadDpadDown    = 626
	Key_GamepadL1          = 627
	Key_GamepadR1          = 628
	Key_GamepadL2          = 629
	Key_GamepadR2          = 630
	Key_GamepadL3          = 631
	Key_GamepadR3          = 632
	Key_GamepadLStickLeft  = 633
	Key_GamepadLStickRight = 634
	Key_GamepadLStickUp    = 635
	Key_GamepadLStickDown  = 636
	Key_GamepadRStickLeft  = 637
	Key_GamepadRStickRight = 638
	Key_GamepadRStickUp    = 639
	Key_GamepadRStickDown  = 640
	Key_ModCtrl            = 641
	Key_ModShift           = 642
	Key_ModAlt             = 643
	Key_ModSuper           = 644
	Key_MouseLeft          = 645
	Key_MouseRight         = 646
	Key_MouseMiddle        = 647
	Key_MouseX1            = 648
	Key_MouseX2            = 649
	Key_MouseWheelX        = 650
	Key_MouseWheelY        = 651
	Key_COUNT              = 652
	Key_NamedKey_BEGIN     = 512
	Key_NamedKey_END       = 652
	Key_NamedKey_COUNT     = 140
	Key_KeysData_SIZE      = 652
	Key_KeysData_OFFSET    = 0
)

// original name: ImGuiLayoutType_
type LayoutType int

const (
	LayoutType_Horizontal = 0
	LayoutType_Vertical   = 1
)

// original name: ImGuiLogType
type LogType int

const (
	LogType_None      = 0
	LogType_TTY       = 1
	LogType_File      = 2
	LogType_Buffer    = 3
	LogType_Clipboard = 4
)

// original name: ImGuiModFlags_
type ModFlags int

const (
	ModFlags_None  = 0
	ModFlags_Ctrl  = 1
	ModFlags_Shift = 2
	ModFlags_Alt   = 4
	ModFlags_Super = 8
	ModFlags_All   = 15
)

// original name: ImGuiMouseButton_
type MouseButton int

const (
	MouseButton_Left   = 0
	MouseButton_Right  = 1
	MouseButton_Middle = 2
	MouseButton_COUNT  = 5
)

// original name: ImGuiMouseCursor_
type MouseCursor int

const (
	MouseCursor_None       = -1
	MouseCursor_Arrow      = 0
	MouseCursor_TextInput  = 1
	MouseCursor_ResizeAll  = 2
	MouseCursor_ResizeNS   = 3
	MouseCursor_ResizeEW   = 4
	MouseCursor_ResizeNESW = 5
	MouseCursor_ResizeNWSE = 6
	MouseCursor_Hand       = 7
	MouseCursor_NotAllowed = 8
	MouseCursor_COUNT      = 9
)

// original name: ImGuiNavHighlightFlags_
type NavHighlightFlags int

const (
	NavHighlightFlags_None        = 0
	NavHighlightFlags_TypeDefault = 1
	NavHighlightFlags_TypeThin    = 2
	NavHighlightFlags_AlwaysDraw  = 4
	NavHighlightFlags_NoRounding  = 8
)

// original name: ImGuiNavInput
type NavInput int

const (
	NavInput_Activate    = 0
	NavInput_Cancel      = 1
	NavInput_Input       = 2
	NavInput_Menu        = 3
	NavInput_DpadLeft    = 4
	NavInput_DpadRight   = 5
	NavInput_DpadUp      = 6
	NavInput_DpadDown    = 7
	NavInput_LStickLeft  = 8
	NavInput_LStickRight = 9
	NavInput_LStickUp    = 10
	NavInput_LStickDown  = 11
	NavInput_FocusPrev   = 12
	NavInput_FocusNext   = 13
	NavInput_TweakSlow   = 14
	NavInput_TweakFast   = 15
	NavInput_COUNT       = 16
)

// original name: ImGuiNavLayer
type NavLayer int

const (
	NavLayer_Main  = 0
	NavLayer_Menu  = 1
	NavLayer_COUNT = 2
)

// original name: ImGuiNavMoveFlags_
type NavMoveFlags int

const (
	NavMoveFlags_None                = 0
	NavMoveFlags_LoopX               = 1
	NavMoveFlags_LoopY               = 2
	NavMoveFlags_WrapX               = 4
	NavMoveFlags_WrapY               = 8
	NavMoveFlags_AllowCurrentNavId   = 16
	NavMoveFlags_AlsoScoreVisibleSet = 32
	NavMoveFlags_ScrollToEdgeY       = 64
	NavMoveFlags_Forwarded           = 128
	NavMoveFlags_DebugNoResult       = 256
	NavMoveFlags_FocusApi            = 512
	NavMoveFlags_Tabbing             = 1024
	NavMoveFlags_Activate            = 2048
	NavMoveFlags_DontSetNavHighlight = 4096
)

// original name: ImGuiNextItemDataFlags_
type NextItemDataFlags int

const (
	NextItemDataFlags_None     = 0
	NextItemDataFlags_HasWidth = 1
	NextItemDataFlags_HasOpen  = 2
)

// original name: ImGuiNextWindowDataFlags_
type NextWindowDataFlags int

const (
	NextWindowDataFlags_None              = 0
	NextWindowDataFlags_HasPos            = 1
	NextWindowDataFlags_HasSize           = 2
	NextWindowDataFlags_HasContentSize    = 4
	NextWindowDataFlags_HasCollapsed      = 8
	NextWindowDataFlags_HasSizeConstraint = 16
	NextWindowDataFlags_HasFocus          = 32
	NextWindowDataFlags_HasBgAlpha        = 64
	NextWindowDataFlags_HasScroll         = 128
	NextWindowDataFlags_HasViewport       = 256
	NextWindowDataFlags_HasDock           = 512
	NextWindowDataFlags_HasWindowClass    = 1024
)

// original name: ImGuiOldColumnFlags_
type OldColumnFlags int

const (
	OldColumnFlags_None                   = 0
	OldColumnFlags_NoBorder               = 1
	OldColumnFlags_NoResize               = 2
	OldColumnFlags_NoPreserveWidths       = 4
	OldColumnFlags_NoForceWithinWindow    = 8
	OldColumnFlags_GrowParentContentsSize = 16
)

// original name: ImGuiPlotType
type PlotType int

const (
	PlotType_Lines     = 0
	PlotType_Histogram = 1
)

// original name: ImGuiPopupFlags_
type PopupFlags int

const (
	PopupFlags_None                    = 0
	PopupFlags_MouseButtonLeft         = 0
	PopupFlags_MouseButtonRight        = 1
	PopupFlags_MouseButtonMiddle       = 2
	PopupFlags_MouseButtonMask_        = 31
	PopupFlags_MouseButtonDefault_     = 1
	PopupFlags_NoOpenOverExistingPopup = 32
	PopupFlags_NoOpenOverItems         = 64
	PopupFlags_AnyPopupId              = 128
	PopupFlags_AnyPopupLevel           = 256
	PopupFlags_AnyPopup                = 384
)

// original name: ImGuiPopupPositionPolicy
type PopupPositionPolicy int

const (
	PopupPositionPolicy_Default  = 0
	PopupPositionPolicy_ComboBox = 1
	PopupPositionPolicy_Tooltip  = 2
)

// original name: ImGuiScrollFlags_
type ScrollFlags int

const (
	ScrollFlags_None               = 0
	ScrollFlags_KeepVisibleEdgeX   = 1
	ScrollFlags_KeepVisibleEdgeY   = 2
	ScrollFlags_KeepVisibleCenterX = 4
	ScrollFlags_KeepVisibleCenterY = 8
	ScrollFlags_AlwaysCenterX      = 16
	ScrollFlags_AlwaysCenterY      = 32
	ScrollFlags_NoScrollParent     = 64
	ScrollFlags_MaskX_             = 21
	ScrollFlags_MaskY_             = 42
)

// original name: ImGuiSelectableFlagsPrivate_
type SelectableFlagsPrivate int

const (
	SelectableFlags_NoHoldingActiveID    = 1048576
	SelectableFlags_SelectOnNav          = 2097152
	SelectableFlags_SelectOnClick        = 4194304
	SelectableFlags_SelectOnRelease      = 8388608
	SelectableFlags_SpanAvailWidth       = 16777216
	SelectableFlags_DrawHoveredWhenHeld  = 33554432
	SelectableFlags_SetNavIdOnHover      = 67108864
	SelectableFlags_NoPadWithHalfSpacing = 134217728
)

// original name: ImGuiSelectableFlags_
type SelectableFlags int

const (
	SelectableFlags_None             = 0
	SelectableFlags_DontClosePopups  = 1
	SelectableFlags_SpanAllColumns   = 2
	SelectableFlags_AllowDoubleClick = 4
	SelectableFlags_Disabled         = 8
	SelectableFlags_AllowItemOverlap = 16
)

// original name: ImGuiSeparatorFlags_
type SeparatorFlags int

const (
	SeparatorFlags_None           = 0
	SeparatorFlags_Horizontal     = 1
	SeparatorFlags_Vertical       = 2
	SeparatorFlags_SpanAllColumns = 4
)

// original name: ImGuiSliderFlagsPrivate_
type SliderFlagsPrivate int

const (
	SliderFlags_Vertical = 1048576
	SliderFlags_ReadOnly = 2097152
)

// original name: ImGuiSliderFlags_
type SliderFlags int

const (
	SliderFlags_None            = 0
	SliderFlags_AlwaysClamp     = 16
	SliderFlags_Logarithmic     = 32
	SliderFlags_NoRoundToFormat = 64
	SliderFlags_NoInput         = 128
	SliderFlags_InvalidMask_    = 1879048207
)

// original name: ImGuiSortDirection_
type SortDirection int

const (
	SortDirection_None       = 0
	SortDirection_Ascending  = 1
	SortDirection_Descending = 2
)

// original name: ImGuiStyleVar_
type StyleVar int

const (
	StyleVar_Alpha               = 0
	StyleVar_DisabledAlpha       = 1
	StyleVar_WindowPadding       = 2
	StyleVar_WindowRounding      = 3
	StyleVar_WindowBorderSize    = 4
	StyleVar_WindowMinSize       = 5
	StyleVar_WindowTitleAlign    = 6
	StyleVar_ChildRounding       = 7
	StyleVar_ChildBorderSize     = 8
	StyleVar_PopupRounding       = 9
	StyleVar_PopupBorderSize     = 10
	StyleVar_FramePadding        = 11
	StyleVar_FrameRounding       = 12
	StyleVar_FrameBorderSize     = 13
	StyleVar_ItemSpacing         = 14
	StyleVar_ItemInnerSpacing    = 15
	StyleVar_IndentSpacing       = 16
	StyleVar_CellPadding         = 17
	StyleVar_ScrollbarSize       = 18
	StyleVar_ScrollbarRounding   = 19
	StyleVar_GrabMinSize         = 20
	StyleVar_GrabRounding        = 21
	StyleVar_TabRounding         = 22
	StyleVar_ButtonTextAlign     = 23
	StyleVar_SelectableTextAlign = 24
	StyleVar_COUNT               = 25
)

// original name: ImGuiTabBarFlagsPrivate_
type TabBarFlagsPrivate int

const (
	TabBarFlags_DockNode     = 1048576
	TabBarFlags_IsFocused    = 2097152
	TabBarFlags_SaveSettings = 4194304
)

// original name: ImGuiTabBarFlags_
type TabBarFlags int

const (
	TabBarFlags_None                         = 0
	TabBarFlags_Reorderable                  = 1
	TabBarFlags_AutoSelectNewTabs            = 2
	TabBarFlags_TabListPopupButton           = 4
	TabBarFlags_NoCloseWithMiddleMouseButton = 8
	TabBarFlags_NoTabListScrollingButtons    = 16
	TabBarFlags_NoTooltip                    = 32
	TabBarFlags_FittingPolicyResizeDown      = 64
	TabBarFlags_FittingPolicyScroll          = 128
	TabBarFlags_FittingPolicyMask_           = 192
	TabBarFlags_FittingPolicyDefault_        = 64
)

// original name: ImGuiTabItemFlagsPrivate_
type TabItemFlagsPrivate int

const (
	TabItemFlags_SectionMask_  = 192
	TabItemFlags_NoCloseButton = 1048576
	TabItemFlags_Button        = 2097152
	TabItemFlags_Unsorted      = 4194304
	TabItemFlags_Preview       = 8388608
)

// original name: ImGuiTabItemFlags_
type TabItemFlags int

const (
	TabItemFlags_None                         = 0
	TabItemFlags_UnsavedDocument              = 1
	TabItemFlags_SetSelected                  = 2
	TabItemFlags_NoCloseWithMiddleMouseButton = 4
	TabItemFlags_NoPushId                     = 8
	TabItemFlags_NoTooltip                    = 16
	TabItemFlags_NoReorder                    = 32
	TabItemFlags_Leading                      = 64
	TabItemFlags_Trailing                     = 128
)

// original name: ImGuiTableBgTarget_
type TableBgTarget int

const (
	TableBgTarget_None   = 0
	TableBgTarget_RowBg0 = 1
	TableBgTarget_RowBg1 = 2
	TableBgTarget_CellBg = 3
)

// original name: ImGuiTableColumnFlags_
type TableColumnFlags int

const (
	TableColumnFlags_None                 = 0
	TableColumnFlags_Disabled             = 1
	TableColumnFlags_DefaultHide          = 2
	TableColumnFlags_DefaultSort          = 4
	TableColumnFlags_WidthStretch         = 8
	TableColumnFlags_WidthFixed           = 16
	TableColumnFlags_NoResize             = 32
	TableColumnFlags_NoReorder            = 64
	TableColumnFlags_NoHide               = 128
	TableColumnFlags_NoClip               = 256
	TableColumnFlags_NoSort               = 512
	TableColumnFlags_NoSortAscending      = 1024
	TableColumnFlags_NoSortDescending     = 2048
	TableColumnFlags_NoHeaderLabel        = 4096
	TableColumnFlags_NoHeaderWidth        = 8192
	TableColumnFlags_PreferSortAscending  = 16384
	TableColumnFlags_PreferSortDescending = 32768
	TableColumnFlags_IndentEnable         = 65536
	TableColumnFlags_IndentDisable        = 131072
	TableColumnFlags_IsEnabled            = 16777216
	TableColumnFlags_IsVisible            = 33554432
	TableColumnFlags_IsSorted             = 67108864
	TableColumnFlags_IsHovered            = 134217728
	TableColumnFlags_WidthMask_           = 24
	TableColumnFlags_IndentMask_          = 196608
	TableColumnFlags_StatusMask_          = 251658240
	TableColumnFlags_NoDirectResize_      = 1073741824
)

// original name: ImGuiTableFlags_
type TableFlags int

const (
	TableFlags_None                       = 0
	TableFlags_Resizable                  = 1
	TableFlags_Reorderable                = 2
	TableFlags_Hideable                   = 4
	TableFlags_Sortable                   = 8
	TableFlags_NoSavedSettings            = 16
	TableFlags_ContextMenuInBody          = 32
	TableFlags_RowBg                      = 64
	TableFlags_BordersInnerH              = 128
	TableFlags_BordersOuterH              = 256
	TableFlags_BordersInnerV              = 512
	TableFlags_BordersOuterV              = 1024
	TableFlags_BordersH                   = 384
	TableFlags_BordersV                   = 1536
	TableFlags_BordersInner               = 640
	TableFlags_BordersOuter               = 1280
	TableFlags_Borders                    = 1920
	TableFlags_NoBordersInBody            = 2048
	TableFlags_NoBordersInBodyUntilResize = 4096
	TableFlags_SizingFixedFit             = 8192
	TableFlags_SizingFixedSame            = 16384
	TableFlags_SizingStretchProp          = 24576
	TableFlags_SizingStretchSame          = 32768
	TableFlags_NoHostExtendX              = 65536
	TableFlags_NoHostExtendY              = 131072
	TableFlags_NoKeepColumnsVisible       = 262144
	TableFlags_PreciseWidths              = 524288
	TableFlags_NoClip                     = 1048576
	TableFlags_PadOuterX                  = 2097152
	TableFlags_NoPadOuterX                = 4194304
	TableFlags_NoPadInnerX                = 8388608
	TableFlags_ScrollX                    = 16777216
	TableFlags_ScrollY                    = 33554432
	TableFlags_SortMulti                  = 67108864
	TableFlags_SortTristate               = 134217728
	TableFlags_SizingMask_                = 57344
)

// original name: ImGuiTableRowFlags_
type TableRowFlags int

const (
	TableRowFlags_None    = 0
	TableRowFlags_Headers = 1
)

// original name: ImGuiTextFlags_
type TextFlags int

const (
	TextFlags_None                       = 0
	TextFlags_NoWidthForLargeClippedText = 1
)

// original name: ImGuiTooltipFlags_
type TooltipFlags int

const (
	TooltipFlags_None                    = 0
	TooltipFlags_OverridePreviousTooltip = 1
)

// original name: ImGuiTreeNodeFlagsPrivate_
type TreeNodeFlagsPrivate int

const (
	TreeNodeFlags_ClipLabelForTrailingButton = 1048576
)

// original name: ImGuiTreeNodeFlags_
type TreeNodeFlags int

const (
	TreeNodeFlags_None                 = 0
	TreeNodeFlags_Selected             = 1
	TreeNodeFlags_Framed               = 2
	TreeNodeFlags_AllowItemOverlap     = 4
	TreeNodeFlags_NoTreePushOnOpen     = 8
	TreeNodeFlags_NoAutoOpenOnLog      = 16
	TreeNodeFlags_DefaultOpen          = 32
	TreeNodeFlags_OpenOnDoubleClick    = 64
	TreeNodeFlags_OpenOnArrow          = 128
	TreeNodeFlags_Leaf                 = 256
	TreeNodeFlags_Bullet               = 512
	TreeNodeFlags_FramePadding         = 1024
	TreeNodeFlags_SpanAvailWidth       = 2048
	TreeNodeFlags_SpanFullWidth        = 4096
	TreeNodeFlags_NavLeftJumpsBackHere = 8192
	TreeNodeFlags_CollapsingHeader     = 26
)

// original name: ImGuiViewportFlags_
type ViewportFlags int

const (
	ViewportFlags_None                = 0
	ViewportFlags_IsPlatformWindow    = 1
	ViewportFlags_IsPlatformMonitor   = 2
	ViewportFlags_OwnedByApp          = 4
	ViewportFlags_NoDecoration        = 8
	ViewportFlags_NoTaskBarIcon       = 16
	ViewportFlags_NoFocusOnAppearing  = 32
	ViewportFlags_NoFocusOnClick      = 64
	ViewportFlags_NoInputs            = 128
	ViewportFlags_NoRendererClear     = 256
	ViewportFlags_TopMost             = 512
	ViewportFlags_Minimized           = 1024
	ViewportFlags_NoAutoMerge         = 2048
	ViewportFlags_CanHostOtherWindows = 4096
)

// original name: ImGuiWindowDockStyleCol
type WindowDockStyleCol int

const (
	WindowDockStyleCol_Text               = 0
	WindowDockStyleCol_Tab                = 1
	WindowDockStyleCol_TabHovered         = 2
	WindowDockStyleCol_TabActive          = 3
	WindowDockStyleCol_TabUnfocused       = 4
	WindowDockStyleCol_TabUnfocusedActive = 5
	WindowDockStyleCol_COUNT              = 6
)

// original name: ImGuiWindowFlags_
type WindowFlags int

const (
	WindowFlags_None                      = 0
	WindowFlags_NoTitleBar                = 1
	WindowFlags_NoResize                  = 2
	WindowFlags_NoMove                    = 4
	WindowFlags_NoScrollbar               = 8
	WindowFlags_NoScrollWithMouse         = 16
	WindowFlags_NoCollapse                = 32
	WindowFlags_AlwaysAutoResize          = 64
	WindowFlags_NoBackground              = 128
	WindowFlags_NoSavedSettings           = 256
	WindowFlags_NoMouseInputs             = 512
	WindowFlags_MenuBar                   = 1024
	WindowFlags_HorizontalScrollbar       = 2048
	WindowFlags_NoFocusOnAppearing        = 4096
	WindowFlags_NoBringToFrontOnFocus     = 8192
	WindowFlags_AlwaysVerticalScrollbar   = 16384
	WindowFlags_AlwaysHorizontalScrollbar = 32768
	WindowFlags_AlwaysUseWindowPadding    = 65536
	WindowFlags_NoNavInputs               = 262144
	WindowFlags_NoNavFocus                = 524288
	WindowFlags_UnsavedDocument           = 1048576
	WindowFlags_NoDocking                 = 2097152
	WindowFlags_NoNav                     = 786432
	WindowFlags_NoDecoration              = 43
	WindowFlags_NoInputs                  = 786944
	WindowFlags_NavFlattened              = 8388608
	WindowFlags_ChildWindow               = 16777216
	WindowFlags_Tooltip                   = 33554432
	WindowFlags_Popup                     = 67108864
	WindowFlags_Modal                     = 134217728
	WindowFlags_ChildMenu                 = 268435456
	WindowFlags_DockNodeHost              = 536870912
)
