/********************************************************************************
** Form generated from reading UI file 'csystemmain.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_CSYSTEMMAIN_H
#define UI_CSYSTEMMAIN_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QGroupBox>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QListView>
#include <QtWidgets/QMainWindow>
#include <QtWidgets/QMenuBar>
#include <QtWidgets/QStackedWidget>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_CSystemMain
{
public:
    QWidget *centralwidget;
    QWidget *layoutWidget;
    QHBoxLayout *horizontalLayout;
    QLabel *label;
    QLineEdit *lineEdit;
    QWidget *widget;
    QStackedWidget *bottombar;
    QWidget *page1;
    QHBoxLayout *horizontalLayout_3;
    QGroupBox *groupBox;
    QGroupBox *groupBox_3;
    QGroupBox *groupBox_2;
    QWidget *page2;
    QGroupBox *groupBox_4;
    QWidget *page3;
    QGroupBox *groupBox_5;
    QListView *side_menu;
    QMenuBar *menubar;

    void setupUi(QMainWindow *CSystemMain)
    {
        if (CSystemMain->objectName().isEmpty())
            CSystemMain->setObjectName(QString::fromUtf8("CSystemMain"));
        CSystemMain->resize(1920, 1047);
        QSizePolicy sizePolicy(QSizePolicy::Expanding, QSizePolicy::Expanding);
        sizePolicy.setHorizontalStretch(0);
        sizePolicy.setVerticalStretch(0);
        sizePolicy.setHeightForWidth(CSystemMain->sizePolicy().hasHeightForWidth());
        CSystemMain->setSizePolicy(sizePolicy);
        centralwidget = new QWidget(CSystemMain);
        centralwidget->setObjectName(QString::fromUtf8("centralwidget"));
        layoutWidget = new QWidget(centralwidget);
        layoutWidget->setObjectName(QString::fromUtf8("layoutWidget"));
        layoutWidget->setGeometry(QRect(-2, 10, 221, 30));
        horizontalLayout = new QHBoxLayout(layoutWidget);
        horizontalLayout->setObjectName(QString::fromUtf8("horizontalLayout"));
        horizontalLayout->setContentsMargins(0, 0, 0, 0);
        label = new QLabel(layoutWidget);
        label->setObjectName(QString::fromUtf8("label"));
        QFont font;
        font.setPointSize(13);
        font.setBold(false);
        font.setWeight(50);
        label->setFont(font);

        horizontalLayout->addWidget(label);

        lineEdit = new QLineEdit(layoutWidget);
        lineEdit->setObjectName(QString::fromUtf8("lineEdit"));

        horizontalLayout->addWidget(lineEdit);

        widget = new QWidget(centralwidget);
        widget->setObjectName(QString::fromUtf8("widget"));
        widget->setGeometry(QRect(220, 0, 1700, 900));
        QSizePolicy sizePolicy1(QSizePolicy::Preferred, QSizePolicy::Expanding);
        sizePolicy1.setHorizontalStretch(0);
        sizePolicy1.setVerticalStretch(0);
        sizePolicy1.setHeightForWidth(widget->sizePolicy().hasHeightForWidth());
        widget->setSizePolicy(sizePolicy1);
        bottombar = new QStackedWidget(centralwidget);
        bottombar->setObjectName(QString::fromUtf8("bottombar"));
        bottombar->setGeometry(QRect(250, 900, 1641, 122));
        page1 = new QWidget();
        page1->setObjectName(QString::fromUtf8("page1"));
        page1->setMaximumSize(QSize(1641, 118));
        horizontalLayout_3 = new QHBoxLayout(page1);
        horizontalLayout_3->setObjectName(QString::fromUtf8("horizontalLayout_3"));
        groupBox = new QGroupBox(page1);
        groupBox->setObjectName(QString::fromUtf8("groupBox"));
        groupBox->setMinimumSize(QSize(820, 111));

        horizontalLayout_3->addWidget(groupBox);

        groupBox_3 = new QGroupBox(page1);
        groupBox_3->setObjectName(QString::fromUtf8("groupBox_3"));
        groupBox_3->setMinimumSize(QSize(0, 111));

        horizontalLayout_3->addWidget(groupBox_3);

        groupBox_2 = new QGroupBox(page1);
        groupBox_2->setObjectName(QString::fromUtf8("groupBox_2"));
        groupBox_2->setMinimumSize(QSize(0, 111));

        horizontalLayout_3->addWidget(groupBox_2);

        bottombar->addWidget(page1);
        page2 = new QWidget();
        page2->setObjectName(QString::fromUtf8("page2"));
        groupBox_4 = new QGroupBox(page2);
        groupBox_4->setObjectName(QString::fromUtf8("groupBox_4"));
        groupBox_4->setGeometry(QRect(0, 0, 1641, 111));
        bottombar->addWidget(page2);
        page3 = new QWidget();
        page3->setObjectName(QString::fromUtf8("page3"));
        groupBox_5 = new QGroupBox(page3);
        groupBox_5->setObjectName(QString::fromUtf8("groupBox_5"));
        groupBox_5->setGeometry(QRect(0, 0, 1641, 111));
        bottombar->addWidget(page3);
        side_menu = new QListView(centralwidget);
        side_menu->setObjectName(QString::fromUtf8("side_menu"));
        side_menu->setGeometry(QRect(0, 40, 221, 981));
        CSystemMain->setCentralWidget(centralwidget);
        menubar = new QMenuBar(CSystemMain);
        menubar->setObjectName(QString::fromUtf8("menubar"));
        menubar->setGeometry(QRect(0, 0, 1920, 26));
        CSystemMain->setMenuBar(menubar);

        retranslateUi(CSystemMain);

        bottombar->setCurrentIndex(0);


        QMetaObject::connectSlotsByName(CSystemMain);
    } // setupUi

    void retranslateUi(QMainWindow *CSystemMain)
    {
        CSystemMain->setWindowTitle(QCoreApplication::translate("CSystemMain", "MainWindow", nullptr));
        label->setText(QCoreApplication::translate("CSystemMain", "   \346\220\234\347\264\242  ", nullptr));
        groupBox->setTitle(QCoreApplication::translate("CSystemMain", "\345\205\250\345\261\200\344\277\241\346\201\257", nullptr));
        groupBox_3->setTitle(QCoreApplication::translate("CSystemMain", "\346\217\220\351\206\222", nullptr));
        groupBox_2->setTitle(QCoreApplication::translate("CSystemMain", "\350\255\246\346\212\245", nullptr));
        groupBox_4->setTitle(QCoreApplication::translate("CSystemMain", "\345\212\237\350\203\275\350\257\264\346\230\216", nullptr));
        groupBox_5->setTitle(QCoreApplication::translate("CSystemMain", "\347\263\273\347\273\237\350\255\246\345\221\212", nullptr));
    } // retranslateUi

};

namespace Ui {
    class CSystemMain: public Ui_CSystemMain {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_CSYSTEMMAIN_H
