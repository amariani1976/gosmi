package smi

import (
	"fmt"
	"unsafe"

	"github.com/amariani1976/gosmi/smi/internal"
	"github.com/amariani1976/gosmi/types"
)

// char *smiLoadModule(const char *module)
func LoadModule(module string) string {
	checkInit()
	modulePtr, err := internal.GetModule(module)
	if err != nil {
		fmt.Println(err)
	}
	if modulePtr == nil {
		return ""
	}
	return modulePtr.Name.String()
}

// int smiIsLoaded(const char *module)
func IsLoaded(module string) bool {
	checkInit()
	return internal.FindModuleByName(module) != nil
}

// SmiModule *smiGetModule(const char *module)
func GetModule(module string) *types.SmiModule {
	if module == "" {
		return nil
	}
	modulePtr, _ := internal.GetModule(module)
	if modulePtr == nil {
		return nil
	}
	return &modulePtr.SmiModule
}

// SmiModule *smiGetFirstModule(void)
func GetFirstModule() *types.SmiModule {
	modulePtr := internal.GetFirstModule()
	if modulePtr == nil {
		return nil
	}
	return &modulePtr.SmiModule
}

// SmiModule *smiGetNextModule(SmiModule *smiModulePtr)
func GetNextModule(smiModulePtr *types.SmiModule) *types.SmiModule {
	if smiModulePtr == nil {
		return nil
	}
	modulePtr := (*internal.Module)(unsafe.Pointer(smiModulePtr))
	if modulePtr.Next == nil {
		return nil
	}
	return &modulePtr.Next.SmiModule

}

// SmiNode *smiGetModuleIdentityNode(SmiModule *smiModulePtr)
func GetModuleIdentityNode(smiModulePtr *types.SmiModule) *types.SmiNode {
	if smiModulePtr == nil {
		return nil
	}
	modulePtr := (*internal.Module)(unsafe.Pointer(smiModulePtr))
	if modulePtr.Identity == nil {
		return nil
	}
	return modulePtr.Identity.GetSmiNode()
}

// GetAllNodes restituisce tutti i nodi dichiarati OBJECT-TYPE e NOTIFICATION-TYPE
func GetAllNodes(module *types.SmiModule) []*types.SmiNode {
    var nodes []*types.SmiNode
    for symbol := module.FirstSymbol; symbol != nil; symbol = symbol.Next {
        node, ok := symbol.Value.(*types.SmiNode)
        if ok && (node.Decl == types.SMI_DECL_OBJECTTYPE || node.Decl == types.SMI_DECL_NOTIFICATIONTYPE) {
            nodes = append(nodes, node)
        }
    }
    return nodes
}