#ifndef CLIPBOARDHELPER_H
#define CLIPBOARDHELPER_H

#include <QApplication>
#include <QClipboard>
#include <QObject>
#include <QDebug>

class ClipboardHelper : public QObject
{
    Q_OBJECT
public:
    explicit ClipboardHelper(QObject *parent = 0) : QObject(parent) {
        clipboard = QApplication::clipboard();
    }

    Q_INVOKABLE void setText(QString text){
        clipboard->setText(text, QClipboard::Clipboard);
    }

    Q_INVOKABLE QVariant getText(){
       return clipboard->text();
    }

private:
    QClipboard *clipboard;
};

#endif // CLIPBOARDHELPER_H
