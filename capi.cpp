#include <QApplication>
#include <QClipboard>
#include "capi.h"

QClipboard_ *getClipboard() {
	return QApplication::clipboard();
}

const char *getText(QClipboard_ *clip) {
	return reinterpret_cast<QClipboard*>(clip)->text().toStdString().c_str();
}

void setText(QClipboard_ *clip, const char *text) {
	reinterpret_cast<QClipboard*>(clip)->setText(QString(text));
}