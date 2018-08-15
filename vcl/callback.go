package vcl

import (
	"unsafe"

	. "github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/rtl"
	. "github.com/ying32/govcl/vcl/types"
)

// 回调过程
func callbackProc(f uintptr, args uintptr, argcount int) uintptr {
	v, ok := CallbackMap[f]
	if ok {

		getVal := func(i int) uintptr {
			return DGetParam(i, args).Value
		}

		switch v.(type) {
		// func(sender IObject)
		case TNotifyEvent:
			v.(TNotifyEvent)(
				ObjectFromInst(getVal(0)))

		// func(sender IObject, button TUDBtnType)
		case TUDClickEvent:
			v.(TUDClickEvent)(
				ObjectFromInst(getVal(0)),
				TUDBtnType(getVal(1)))

		// func(sender IObject, item *TListItem, change int32)
		case TLVChangeEvent:
			v.(TLVChangeEvent)(
				ObjectFromInst(getVal(0)),
				ListItemFromInst(getVal(1)),
				TItemChange(getVal(2)))

		// func(sender IObject, action *TCloseAction) // Action *uintptr
		case TCloseEvent:
			v.(TCloseEvent)(
				ObjectFromInst(getVal(0)),
				(*TCloseAction)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, canClose *bool) //CanClose *uintptr
		case TCloseQueryEvent:
			v.(TCloseQueryEvent)(
				ObjectFromInst(getVal(0)),
				(*bool)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, source *TMenuItem, rebuild bool)
		case TMenuChangeEvent:
			v.(TMenuChangeEvent)(
				ObjectFromInst(getVal(0)),
				MenuItemFromInst(getVal(1)),
				DBoolToGoBool(getVal(2)))

		// func(sender IObject, node *TreeNode)
		case TTVChangedEvent:
			v.(TTVChangedEvent)(
				ObjectFromInst(getVal(0)),
				TreeNodeFromInst(getVal(1)))

		// func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType
		case TSysLinkEvent:
			v.(TSysLinkEvent)(
				ObjectFromInst(getVal(0)),
				DStrToGoStr(getVal(1)),
				TSysLinkType(getVal(2)))

		// func(sender, e IObject)
		case TExceptionEvent:
			v.(TExceptionEvent)(
				ObjectFromInst(getVal(0)),
				ExceptionFromInst(getVal(1)))

		// func(sender IObject, key *Char, shift TShiftState)
		case TKeyEvent:

			val := (*Char)(unsafe.Pointer(getVal(1)))
			if rtl.LcLLoaded() {
				*val = Char(uint8(*val))
			}
			v.(TKeyEvent)(ObjectFromInst(getVal(0)), val, TShiftState(getVal(2)))

		// func(sender IObject, key *Char)
		case TKeyPressEvent:

			// 这里要改下，Delphi与FreePascal的区别，Delphi用的Unicode所以用uint16，Lazarus用的UTF8，只有一个字节的。
			val := (*Char)(unsafe.Pointer(getVal(1)))
			if rtl.LcLLoaded() {
				*val = Char(uint8(*val))
			}
			v.(TKeyPressEvent)(ObjectFromInst(getVal(0)), val)

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
		case TMouseEvent:
			v.(TMouseEvent)(
				ObjectFromInst(getVal(0)),
				TMouseButton(getVal(1)),
				TShiftState(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)))

		// func(sender IObject, shift TShiftState, x, y int32)
		case TMouseMoveEvent:
			v.(TMouseMoveEvent)(
				ObjectFromInst(getVal(0)),
				TShiftState(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

		// func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
		case TMouseWheelEvent:
			v.(TMouseWheelEvent)(
				ObjectFromInst(getVal(0)),
				TShiftState(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

			// func(control IWinControl, index int32, aRect TRect, state TOwnerDrawState)
		case TDrawItemEvent:
			v.(TDrawItemEvent)(
				WinControlFromInst(getVal(0)),
				int32(getVal(1)),
				*(*TRect)(unsafe.Pointer(getVal(2))),
				TOwnerDrawState(getVal(3)))

			// func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)
		case TMenuDrawItemEvent:
			v.(TMenuDrawItemEvent)(
				ObjectFromInst(getVal(0)),
				CanvasFromInst(getVal(1)),
				*(*TRect)(unsafe.Pointer(getVal(2))),
				DBoolToGoBool(getVal(3)))

			// type TLVNotifyEvent func(sender IObject, item *TListItem)
		case TLVNotifyEvent:
			v.(TLVNotifyEvent)(
				ObjectFromInst(getVal(0)),
				ListItemFromInst(getVal(1)))

			// type TLVColumnClickEvent func(sender IObject, column *TListColumn)
		case TLVColumnClickEvent:
			v.(TLVColumnClickEvent)(
				ObjectFromInst(getVal(0)),
				ListColumnFromInst(getVal(1)))

			// type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)
		case TLVColumnRClickEvent:
			v.(TLVColumnRClickEvent)(
				ObjectFromInst(getVal(0)),
				ListColumnFromInst(getVal(1)),
				TPoint{X: int32(getVal(2)), Y: int32(getVal(3))})

			// type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)
		case TLVSelectItemEvent:
			v.(TLVSelectItemEvent)(
				ObjectFromInst(getVal(0)),
				ListItemFromInst(getVal(1)),
				DBoolToGoBool(getVal(2)))

			// type TLVCheckedItemEvent func(sender IObject, item *TListItem)
		case TLVCheckedItemEvent:
			v.(TLVCheckedItemEvent)(
				ObjectFromInst(getVal(0)),
				ListItemFromInst(getVal(1)))

			// type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)
		case TTabGetImageEvent:
			v.(TTabGetImageEvent)(
				ObjectFromInst(getVal(0)),
				int32(getVal(1)),
				(*int32)(unsafe.Pointer(getVal(2))))

			// type TTVExpandedEvent func(sender IObject, node *TTreeNode)
		case TTVExpandedEvent:
			v.(TTVExpandedEvent)(
				ObjectFromInst(getVal(0)),
				TreeNodeFromInst(getVal(1)))

		//type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)
		case TLVCompareEvent:
			v.(TLVCompareEvent)(
				ObjectFromInst(getVal(0)),
				ListItemFromInst(getVal(1)),
				ListItemFromInst(getVal(2)),
				int32(getVal(3)),
				(*int32)(unsafe.Pointer(getVal(4))))

		//type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)
		case TTVCompareEvent:
			v.(TTVCompareEvent)(
				ObjectFromInst(getVal(0)),
				TreeNodeFromInst(getVal(1)),
				TreeNodeFromInst(getVal(2)),
				int32(getVal(3)),
				(*int32)(unsafe.Pointer(getVal(4))))

		//type TTVAdvancedCustomDrawEvent func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTVAdvancedCustomDrawEvent:

			v.(TTVAdvancedCustomDrawEvent)(
				TreeViewFromInst(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

		//type TTVAdvancedCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)
		case TTVAdvancedCustomDrawItemEvent:
			v.(TTVAdvancedCustomDrawItemEvent)(
				TreeViewFromInst(getVal(0)),
				TreeNodeFromInst(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				(*bool)(unsafe.Pointer(getVal(4))),
				(*bool)(unsafe.Pointer(getVal(5))))

			//---------------------------------------

		//type TLVAdvancedCustomDrawEvent func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawEvent:
			v.(TLVAdvancedCustomDrawEvent)(
				ListViewFromInst(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

		//type TLVAdvancedCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawItemEvent:
			v.(TLVAdvancedCustomDrawItemEvent)(
				ListViewFromInst(getVal(0)),
				ListItemFromInst(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				(*bool)(unsafe.Pointer(getVal(4))))

		//type TLVAdvancedCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawSubItemEvent:
			v.(TLVAdvancedCustomDrawSubItemEvent)(
				ListViewFromInst(getVal(0)),
				ListItemFromInst(getVal(1)),
				int32(getVal(2)),
				TCustomDrawState(getVal(3)),
				TCustomDrawStage(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

		//-----------------------------
		//type TTBAdvancedCustomDrawEvent func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTBAdvancedCustomDrawEvent:
			v.(TTBAdvancedCustomDrawEvent)(
				ToolBarFromInst(getVal(0)),
				*(*TRect)(unsafe.Pointer(getVal(1))),
				TCustomDrawStage(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

		//type TTBAdvancedCustomDrawBtnEvent func(sender *TToolBar, button *TToolButton, state TCustomDrawState, stage TCustomDrawStage, flags *TTBCustomDrawFlags, defaultDraw *bool)
		case TTBAdvancedCustomDrawBtnEvent:
			v.(TTBAdvancedCustomDrawBtnEvent)(
				ToolBarFromInst(getVal(0)),
				ToolButtonFromInst(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				(*TTBCustomDrawFlags)(unsafe.Pointer(getVal(4))),
				(*bool)(unsafe.Pointer(getVal(5))))

		// TThreadProc
		case TThreadProc:
			v.(TThreadProc)()

		// TDropFilesEvent
		case TDropFilesEvent:
			nLen := int(getVal(2))
			tempArr := make([]string, nLen)
			p := getVal(1)
			for i := 0; i < nLen; i++ {
				tempArr[i] = DGetStringArrOf(p, i)
			}
			v.(TDropFilesEvent)(
				ObjectFromInst(getVal(0)),
				tempArr)

			// TConstrainedResizeEvent
		case TConstrainedResizeEvent:
			v.(TConstrainedResizeEvent)(
				ObjectFromInst(getVal(0)),
				(*int32)(unsafe.Pointer(getVal(1))),
				(*int32)(unsafe.Pointer(getVal(2))),
				(*int32)(unsafe.Pointer(getVal(3))),
				(*int32)(unsafe.Pointer(getVal(4))))

			// func(command uint16, data THelpEventData, callhelp *bool) bool
		case THelpEvent:
			v.(THelpEvent)(
				uint16(getVal(0)),
				THelpEventData(getVal(1)),
				(*bool)(unsafe.Pointer(getVal(2))),
				(*bool)(unsafe.Pointer(getVal(3))))

			// func(msg *TWMKey, handled *bool)
		case TShortCutEvent:
			v.(TShortCutEvent)(
				(*TWMKey)(unsafe.Pointer(getVal(0))),
				(*bool)(unsafe.Pointer(getVal(1))))

			// func(sender IObject, mousePos TPoint, handled *bool)
		case TContextPopupEvent:
			v.(TContextPopupEvent)(
				ObjectFromInst(getVal(0)),
				*(*TPoint)(unsafe.Pointer(getVal(1))),
				(*bool)(unsafe.Pointer(getVal(2))))

			// func(sender, source IObject, x, y int32, state TDragState, accept *bool)
		case TDragOverEvent:
			v.(TDragOverEvent)(
				ObjectFromInst(getVal(0)),
				ObjectFromInst(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				TDragState(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

			//func(sender, source IObject, x, y int32)
		case TDragDropEvent:
			v.(TDragDropEvent)(
				ObjectFromInst(getVal(0)),
				ObjectFromInst(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			//func(sender IObject, dragObject *TDragObject)
		case TStartDragEvent:
			v.(TStartDragEvent)(
				ObjectFromInst(getVal(0)),
				DragObjectFromInst(getVal(1)))

			//func(sender, target IObject, x, y int32)
		case TEndDragEvent:
			v.(TEndDragEvent)(
				ObjectFromInst(getVal(0)),
				ObjectFromInst(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			// func(sender IObject, source *TDragDockObject, x, y int32)
		case TDockDropEvent:
			v.(TDockDropEvent)(
				ObjectFromInst(getVal(0)),
				DragDockObjectFromInst(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			//func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)
		case TDockOverEvent:
			v.(TDockOverEvent)(
				ObjectFromInst(getVal(0)),
				DragDockObjectFromInst(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				TDragState(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

			//func(sender IObject, client *TControl, newTarget *TControl, allow *bool)
		case TUnDockEvent:
			v.(TUnDockEvent)(
				ObjectFromInst(getVal(0)),
				ControlFromInst(getVal(1)),
				ControlFromInst(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//func(sender IObject, dragObject *TDragDockObject)
		case TStartDockEvent:
			v.(TStartDockEvent)(
				ObjectFromInst(getVal(0)),
				DragDockObjectFromInst(getVal(1)))

			//func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)
		case TGetSiteInfoEvent:
			v.(TGetSiteInfoEvent)(
				ObjectFromInst(getVal(0)),
				ControlFromInst(getVal(1)),
				(*TRect)(unsafe.Pointer(getVal(2))),
				*(*TPoint)(unsafe.Pointer(getVal(3))),
				(*bool)(unsafe.Pointer(getVal(4))))

			//func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)
		case TMouseWheelUpDownEvent:
			v.(TMouseWheelUpDownEvent)(
				ObjectFromInst(getVal(0)),
				TShiftState(getVal(1)),
				*(*TPoint)(unsafe.Pointer(getVal(2))),
				(*bool)(unsafe.Pointer(getVal(3))))

			// TMessageEvent = procedure (var Msg: TMsg; var Handled: Boolean) of object;
		case TMessageEvent: // func(msg *TMsg, handled *bool)
			v.(TMessageEvent)(
				(*TMsg)(unsafe.Pointer(getVal(0))),
				(*bool)(unsafe.Pointer(getVal(1))))

			// ---- grid
			//type TMovedEvent func(sender IObject, fromIndex, toIndex int32)
		case TMovedEvent:
			v.(TMovedEvent)(
				ObjectFromInst(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

			//type TDrawCellEvent func(sender IObject, aCol, aRow int32, aRect TRect, state TGridDrawState)
		case TDrawCellEvent:
			v.(TDrawCellEvent)(
				ObjectFromInst(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				*(*TRect)(unsafe.Pointer(getVal(3))),
				TGridDrawState(getVal(4)))

			//type TFixedCellClickEvent func(sender IObject, aCol, aRow int32)
		case TFixedCellClickEvent:
			v.(TFixedCellClickEvent)(
				ObjectFromInst(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

			//type TGetEditEvent func(sender IObject, aCol, aRow int32, value *string)
		case TGetEditEvent:
			str := DStrToGoStr(*(*uintptr)(unsafe.Pointer(getVal(3))))
			v.(TGetEditEvent)(
				ObjectFromInst(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				&str)
			*(*uintptr)(unsafe.Pointer(getVal(3))) = GoStrToDStr(str)

			//type TSelectCellEvent func(sender IObject, aCol, aRow int32, canSelect *bool)
		case TSelectCellEvent:
			v.(TSelectCellEvent)(
				ObjectFromInst(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//type TSetEditEvent func(sender IObject, aCol, aRow int32, value string)
		case TSetEditEvent:
			v.(TSetEditEvent)(
				ObjectFromInst(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				DStrToGoStr(getVal(3)))

			// ---- headercontrol
			//type TDrawSectionEvent func(headerControl *THeaderControl, section *THeaderSection, aRect TRect, pressed bool)
		case TDrawSectionEvent:
			v.(TDrawSectionEvent)(
				HeaderControlFromInst(getVal(0)),
				HeaderSectionFromInst(getVal(1)),
				*(*TRect)(unsafe.Pointer(getVal(2))),
				getVal(3) != 0)

			//type TSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)
		case TSectionNotifyEvent:
			v.(TSectionNotifyEvent)(
				HeaderControlFromInst(getVal(0)),
				HeaderSectionFromInst(getVal(1)))

			//type TSectionTrackEvent func(headerControl *THeaderControl, section *THeaderSection, width int32, state TSectionTrackState)
		case TSectionTrackEvent:
			v.(TSectionTrackEvent)(
				HeaderControlFromInst(getVal(0)),
				HeaderSectionFromInst(getVal(1)),
				int32(getVal(2)),
				TSectionTrackState(getVal(3)))

			//type TSectionDragEvent func(sender IObject, fromSection, toSection *THeaderSection, allowDrag *bool)
		case TSectionDragEvent:
			v.(TSectionDragEvent)(
				ObjectFromInst(getVal(0)),
				HeaderSectionFromInst(getVal(1)),
				HeaderSectionFromInst(getVal(2)),
				(*bool)(unsafe.Pointer(getVal(3))))

			//type TCustomSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)
		case TCustomSectionNotifyEvent:
			v.(TCustomSectionNotifyEvent)(
				HeaderControlFromInst(getVal(0)),
				HeaderSectionFromInst(getVal(1)))

		default:
		}
	}
	return 0
}
