#include "clipboardhelper_moc.h"
#include "clipboard.h"

ClipboardHelper_ *clipboardHelper() {
	return new ClipboardHelper();
}
