/********************************************************************************
** Form generated from reading UI file 'dataview.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_DATAVIEW_H
#define UI_DATAVIEW_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QCheckBox>
#include <QtWidgets/QDateTimeEdit>
#include <QtWidgets/QGridLayout>
#include <QtWidgets/QGroupBox>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QHeaderView>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QListView>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QRadioButton>
#include <QtWidgets/QSpacerItem>
#include <QtWidgets/QTabWidget>
#include <QtWidgets/QTableView>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_DataView
{
public:
    QTabWidget *tabWidget;
    QWidget *tab;
    QTableView *tableView;
    QWidget *tab_2;
    QListView *listView;
    QWidget *tab_3;
    QListView *listView_2;
    QWidget *tab_4;
    QListView *listView_3;
    QWidget *tab_5;
    QListView *listView_4;
    QWidget *tab_6;
    QListView *listView_5;
    QWidget *tab_7;
    QListView *listView_6;
    QWidget *tab_8;
    QTableView *tableView_2;
    QLabel *label;
    QWidget *layoutWidget;
    QVBoxLayout *verticalLayout;
    QGroupBox *GP1;
    QGridLayout *gridLayout;
    QSpacerItem *verticalSpacer_8;
    QLineEdit *lineEdit_2;
    QLabel *label_5;
    QSpacerItem *verticalSpacer_5;
    QLabel *label_4;
    QLabel *label_3;
    QSpacerItem *verticalSpacer;
    QLineEdit *lineEdit_3;
    QHBoxLayout *horizontalLayout;
    QSpacerItem *horizontalSpacer_3;
    QPushButton *pushButton_2;
    QSpacerItem *horizontalSpacer;
    QPushButton *pushButton;
    QSpacerItem *horizontalSpacer_2;
    QLineEdit *lineEdit;
    QSpacerItem *verticalSpacer_10;
    QSpacerItem *verticalSpacer_13;
    QGroupBox *GP2;
    QGridLayout *gridLayout_2;
    QSpacerItem *verticalSpacer_7;
    QHBoxLayout *horizontalLayout_2;
    QSpacerItem *horizontalSpacer_6;
    QPushButton *pushButton_4;
    QSpacerItem *horizontalSpacer_4;
    QPushButton *pushButton_3;
    QSpacerItem *horizontalSpacer_5;
    QDateTimeEdit *dateTimeEdit_2;
    QSpacerItem *verticalSpacer_9;
    QLabel *label_10;
    QLabel *label_6;
    QLabel *label_8;
    QRadioButton *radioButton_2;
    QDateTimeEdit *dateTimeEdit;
    QLabel *label_7;
    QRadioButton *radioButton;
    QSpacerItem *verticalSpacer_3;
    QSpacerItem *verticalSpacer_11;
    QSpacerItem *verticalSpacer_14;
    QGroupBox *GP3;
    QGridLayout *gridLayout_3;
    QCheckBox *checkBox_3;
    QCheckBox *checkBox_4;
    QCheckBox *checkBox_5;
    QSpacerItem *verticalSpacer_2;
    QCheckBox *checkBox_7;
    QRadioButton *radioButton_4;
    QHBoxLayout *horizontalLayout_3;
    QSpacerItem *horizontalSpacer_7;
    QPushButton *pushButton_6;
    QSpacerItem *horizontalSpacer_8;
    QPushButton *pushButton_5;
    QSpacerItem *horizontalSpacer_9;
    QSpacerItem *horizontalSpacer_10;
    QCheckBox *checkBox_6;
    QRadioButton *radioButton_3;
    QSpacerItem *verticalSpacer_4;
    QHBoxLayout *horizontalLayout_4;
    QCheckBox *checkBox;
    QCheckBox *checkBox_2;

    void setupUi(QWidget *DataView)
    {
        if (DataView->objectName().isEmpty())
            DataView->setObjectName(QString::fromUtf8("DataView"));
        DataView->resize(1700, 900);
        tabWidget = new QTabWidget(DataView);
        tabWidget->setObjectName(QString::fromUtf8("tabWidget"));
        tabWidget->setGeometry(QRect(30, 50, 1300, 850));
        tabWidget->setMinimumSize(QSize(1300, 850));
        tabWidget->setStyleSheet(QString::fromUtf8("gridline-color: rgb(255, 255, 255);"));
        tab = new QWidget();
        tab->setObjectName(QString::fromUtf8("tab"));
        tableView = new QTableView(tab);
        tableView->setObjectName(QString::fromUtf8("tableView"));
        tableView->setGeometry(QRect(0, 0, 1293, 814));
        tableView->setStyleSheet(QString::fromUtf8(""));
        tabWidget->addTab(tab, QString());
        tab_2 = new QWidget();
        tab_2->setObjectName(QString::fromUtf8("tab_2"));
        listView = new QListView(tab_2);
        listView->setObjectName(QString::fromUtf8("listView"));
        listView->setGeometry(QRect(0, 0, 1293, 814));
        tabWidget->addTab(tab_2, QString());
        tab_3 = new QWidget();
        tab_3->setObjectName(QString::fromUtf8("tab_3"));
        listView_2 = new QListView(tab_3);
        listView_2->setObjectName(QString::fromUtf8("listView_2"));
        listView_2->setGeometry(QRect(0, 0, 1293, 814));
        tabWidget->addTab(tab_3, QString());
        tab_4 = new QWidget();
        tab_4->setObjectName(QString::fromUtf8("tab_4"));
        listView_3 = new QListView(tab_4);
        listView_3->setObjectName(QString::fromUtf8("listView_3"));
        listView_3->setGeometry(QRect(0, 0, 1293, 814));
        tabWidget->addTab(tab_4, QString());
        tab_5 = new QWidget();
        tab_5->setObjectName(QString::fromUtf8("tab_5"));
        listView_4 = new QListView(tab_5);
        listView_4->setObjectName(QString::fromUtf8("listView_4"));
        listView_4->setGeometry(QRect(0, 0, 1293, 814));
        tabWidget->addTab(tab_5, QString());
        tab_6 = new QWidget();
        tab_6->setObjectName(QString::fromUtf8("tab_6"));
        listView_5 = new QListView(tab_6);
        listView_5->setObjectName(QString::fromUtf8("listView_5"));
        listView_5->setGeometry(QRect(0, 0, 1293, 814));
        tabWidget->addTab(tab_6, QString());
        tab_7 = new QWidget();
        tab_7->setObjectName(QString::fromUtf8("tab_7"));
        listView_6 = new QListView(tab_7);
        listView_6->setObjectName(QString::fromUtf8("listView_6"));
        listView_6->setGeometry(QRect(0, 0, 1293, 814));
        tabWidget->addTab(tab_7, QString());
        tab_8 = new QWidget();
        tab_8->setObjectName(QString::fromUtf8("tab_8"));
        tableView_2 = new QTableView(tab_8);
        tableView_2->setObjectName(QString::fromUtf8("tableView_2"));
        tableView_2->setGeometry(QRect(0, 0, 1293, 814));
        tabWidget->addTab(tab_8, QString());
        label = new QLabel(DataView);
        label->setObjectName(QString::fromUtf8("label"));
        label->setGeometry(QRect(30, 15, 91, 22));
        QFont font;
        font.setPointSize(14);
        font.setBold(true);
        font.setWeight(75);
        label->setFont(font);
        layoutWidget = new QWidget(DataView);
        layoutWidget->setObjectName(QString::fromUtf8("layoutWidget"));
        layoutWidget->setGeometry(QRect(1340, 60, 348, 831));
        verticalLayout = new QVBoxLayout(layoutWidget);
        verticalLayout->setObjectName(QString::fromUtf8("verticalLayout"));
        verticalLayout->setContentsMargins(0, 0, 0, 0);
        GP1 = new QGroupBox(layoutWidget);
        GP1->setObjectName(QString::fromUtf8("GP1"));
        gridLayout = new QGridLayout(GP1);
        gridLayout->setObjectName(QString::fromUtf8("gridLayout"));
        verticalSpacer_8 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer_8, 1, 1, 1, 1);

        lineEdit_2 = new QLineEdit(GP1);
        lineEdit_2->setObjectName(QString::fromUtf8("lineEdit_2"));

        gridLayout->addWidget(lineEdit_2, 2, 3, 1, 1);

        label_5 = new QLabel(GP1);
        label_5->setObjectName(QString::fromUtf8("label_5"));

        gridLayout->addWidget(label_5, 4, 1, 1, 1);

        verticalSpacer_5 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer_5, 0, 0, 1, 1);

        label_4 = new QLabel(GP1);
        label_4->setObjectName(QString::fromUtf8("label_4"));

        gridLayout->addWidget(label_4, 2, 1, 1, 1);

        label_3 = new QLabel(GP1);
        label_3->setObjectName(QString::fromUtf8("label_3"));

        gridLayout->addWidget(label_3, 0, 1, 1, 1);

        verticalSpacer = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer, 5, 1, 1, 1);

        lineEdit_3 = new QLineEdit(GP1);
        lineEdit_3->setObjectName(QString::fromUtf8("lineEdit_3"));

        gridLayout->addWidget(lineEdit_3, 4, 3, 1, 1);

        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setObjectName(QString::fromUtf8("horizontalLayout"));
        horizontalSpacer_3 = new QSpacerItem(13, 17, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout->addItem(horizontalSpacer_3);

        pushButton_2 = new QPushButton(GP1);
        pushButton_2->setObjectName(QString::fromUtf8("pushButton_2"));

        horizontalLayout->addWidget(pushButton_2);

        horizontalSpacer = new QSpacerItem(14, 17, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout->addItem(horizontalSpacer);

        pushButton = new QPushButton(GP1);
        pushButton->setObjectName(QString::fromUtf8("pushButton"));

        horizontalLayout->addWidget(pushButton);

        horizontalSpacer_2 = new QSpacerItem(13, 17, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout->addItem(horizontalSpacer_2);


        gridLayout->addLayout(horizontalLayout, 6, 3, 1, 1);

        lineEdit = new QLineEdit(GP1);
        lineEdit->setObjectName(QString::fromUtf8("lineEdit"));

        gridLayout->addWidget(lineEdit, 0, 3, 1, 1);

        verticalSpacer_10 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer_10, 3, 1, 1, 1);


        verticalLayout->addWidget(GP1);

        verticalSpacer_13 = new QSpacerItem(20, 68, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer_13);

        GP2 = new QGroupBox(layoutWidget);
        GP2->setObjectName(QString::fromUtf8("GP2"));
        gridLayout_2 = new QGridLayout(GP2);
        gridLayout_2->setObjectName(QString::fromUtf8("gridLayout_2"));
        verticalSpacer_7 = new QSpacerItem(57, 15, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_7, 6, 1, 1, 1);

        horizontalLayout_2 = new QHBoxLayout();
        horizontalLayout_2->setObjectName(QString::fromUtf8("horizontalLayout_2"));
        horizontalSpacer_6 = new QSpacerItem(13, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout_2->addItem(horizontalSpacer_6);

        pushButton_4 = new QPushButton(GP2);
        pushButton_4->setObjectName(QString::fromUtf8("pushButton_4"));

        horizontalLayout_2->addWidget(pushButton_4);

        horizontalSpacer_4 = new QSpacerItem(13, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout_2->addItem(horizontalSpacer_4);

        pushButton_3 = new QPushButton(GP2);
        pushButton_3->setObjectName(QString::fromUtf8("pushButton_3"));

        horizontalLayout_2->addWidget(pushButton_3);

        horizontalSpacer_5 = new QSpacerItem(13, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout_2->addItem(horizontalSpacer_5);


        gridLayout_2->addLayout(horizontalLayout_2, 9, 2, 1, 1);

        dateTimeEdit_2 = new QDateTimeEdit(GP2);
        dateTimeEdit_2->setObjectName(QString::fromUtf8("dateTimeEdit_2"));
        dateTimeEdit_2->setDateTime(QDateTime(QDate(2001, 1, 1), QTime(0, 0, 0)));

        gridLayout_2->addWidget(dateTimeEdit_2, 7, 2, 1, 1);

        verticalSpacer_9 = new QSpacerItem(57, 17, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_9, 4, 1, 1, 1);

        label_10 = new QLabel(GP2);
        label_10->setObjectName(QString::fromUtf8("label_10"));

        gridLayout_2->addWidget(label_10, 3, 1, 1, 1);

        label_6 = new QLabel(GP2);
        label_6->setObjectName(QString::fromUtf8("label_6"));

        gridLayout_2->addWidget(label_6, 0, 1, 1, 1);

        label_8 = new QLabel(GP2);
        label_8->setObjectName(QString::fromUtf8("label_8"));

        gridLayout_2->addWidget(label_8, 7, 1, 1, 1);

        radioButton_2 = new QRadioButton(GP2);
        radioButton_2->setObjectName(QString::fromUtf8("radioButton_2"));

        gridLayout_2->addWidget(radioButton_2, 2, 2, 1, 1);

        dateTimeEdit = new QDateTimeEdit(GP2);
        dateTimeEdit->setObjectName(QString::fromUtf8("dateTimeEdit"));

        gridLayout_2->addWidget(dateTimeEdit, 5, 2, 1, 1);

        label_7 = new QLabel(GP2);
        label_7->setObjectName(QString::fromUtf8("label_7"));

        gridLayout_2->addWidget(label_7, 5, 1, 1, 1);

        radioButton = new QRadioButton(GP2);
        radioButton->setObjectName(QString::fromUtf8("radioButton"));

        gridLayout_2->addWidget(radioButton, 1, 2, 1, 1);

        verticalSpacer_3 = new QSpacerItem(20, 15, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_3, 8, 1, 1, 1);

        verticalSpacer_11 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_11, 0, 0, 1, 1);


        verticalLayout->addWidget(GP2);

        verticalSpacer_14 = new QSpacerItem(20, 65, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer_14);

        GP3 = new QGroupBox(layoutWidget);
        GP3->setObjectName(QString::fromUtf8("GP3"));
        gridLayout_3 = new QGridLayout(GP3);
        gridLayout_3->setObjectName(QString::fromUtf8("gridLayout_3"));
        checkBox_3 = new QCheckBox(GP3);
        checkBox_3->setObjectName(QString::fromUtf8("checkBox_3"));

        gridLayout_3->addWidget(checkBox_3, 4, 3, 1, 1);

        checkBox_4 = new QCheckBox(GP3);
        checkBox_4->setObjectName(QString::fromUtf8("checkBox_4"));

        gridLayout_3->addWidget(checkBox_4, 5, 3, 1, 1);

        checkBox_5 = new QCheckBox(GP3);
        checkBox_5->setObjectName(QString::fromUtf8("checkBox_5"));

        gridLayout_3->addWidget(checkBox_5, 6, 3, 1, 1);

        verticalSpacer_2 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_3->addItem(verticalSpacer_2, 0, 0, 2, 1);

        checkBox_7 = new QCheckBox(GP3);
        checkBox_7->setObjectName(QString::fromUtf8("checkBox_7"));

        gridLayout_3->addWidget(checkBox_7, 8, 3, 1, 1);

        radioButton_4 = new QRadioButton(GP3);
        radioButton_4->setObjectName(QString::fromUtf8("radioButton_4"));

        gridLayout_3->addWidget(radioButton_4, 3, 1, 1, 2);

        horizontalLayout_3 = new QHBoxLayout();
        horizontalLayout_3->setObjectName(QString::fromUtf8("horizontalLayout_3"));
        horizontalSpacer_7 = new QSpacerItem(43, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout_3->addItem(horizontalSpacer_7);

        pushButton_6 = new QPushButton(GP3);
        pushButton_6->setObjectName(QString::fromUtf8("pushButton_6"));

        horizontalLayout_3->addWidget(pushButton_6);

        horizontalSpacer_8 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout_3->addItem(horizontalSpacer_8);

        pushButton_5 = new QPushButton(GP3);
        pushButton_5->setObjectName(QString::fromUtf8("pushButton_5"));

        horizontalLayout_3->addWidget(pushButton_5);

        horizontalSpacer_9 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout_3->addItem(horizontalSpacer_9);


        gridLayout_3->addLayout(horizontalLayout_3, 10, 2, 1, 2);

        horizontalSpacer_10 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_3->addItem(horizontalSpacer_10, 10, 0, 1, 2);

        checkBox_6 = new QCheckBox(GP3);
        checkBox_6->setObjectName(QString::fromUtf8("checkBox_6"));

        gridLayout_3->addWidget(checkBox_6, 7, 3, 1, 1);

        radioButton_3 = new QRadioButton(GP3);
        radioButton_3->setObjectName(QString::fromUtf8("radioButton_3"));

        gridLayout_3->addWidget(radioButton_3, 0, 1, 1, 2);

        verticalSpacer_4 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_3->addItem(verticalSpacer_4, 9, 0, 1, 1);

        horizontalLayout_4 = new QHBoxLayout();
        horizontalLayout_4->setObjectName(QString::fromUtf8("horizontalLayout_4"));
        checkBox = new QCheckBox(GP3);
        checkBox->setObjectName(QString::fromUtf8("checkBox"));

        horizontalLayout_4->addWidget(checkBox);

        checkBox_2 = new QCheckBox(GP3);
        checkBox_2->setObjectName(QString::fromUtf8("checkBox_2"));

        horizontalLayout_4->addWidget(checkBox_2);


        gridLayout_3->addLayout(horizontalLayout_4, 2, 3, 1, 1);


        verticalLayout->addWidget(GP3);


        retranslateUi(DataView);

        tabWidget->setCurrentIndex(0);


        QMetaObject::connectSlotsByName(DataView);
    } // setupUi

    void retranslateUi(QWidget *DataView)
    {
        DataView->setWindowTitle(QCoreApplication::translate("DataView", "Form", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab), QCoreApplication::translate("DataView", "\345\205\250\351\203\250", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab_2), QCoreApplication::translate("DataView", "\346\260\224\344\275\223", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab_3), QCoreApplication::translate("DataView", "\346\270\251\345\272\246", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab_4), QCoreApplication::translate("DataView", "PH", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab_5), QCoreApplication::translate("DataView", "\347\224\265\345\257\274\347\216\207", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab_6), QCoreApplication::translate("DataView", "\346\260\247\345\220\253\351\207\217", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab_7), QCoreApplication::translate("DataView", "\351\207\215\351\207\221\345\261\236", nullptr));
        tabWidget->setTabText(tabWidget->indexOf(tab_8), QCoreApplication::translate("DataView", "\345\275\225\345\205\245\346\225\260\346\215\256", nullptr));
        label->setText(QCoreApplication::translate("DataView", "\346\225\260\346\215\256\346\265\217\350\247\210", nullptr));
        GP1->setTitle(QCoreApplication::translate("DataView", "\346\220\234\347\264\242", nullptr));
        label_5->setText(QCoreApplication::translate("DataView", "\346\220\234\347\264\242\346\211\200\345\261\236", nullptr));
        label_4->setText(QCoreApplication::translate("DataView", "\346\220\234\347\264\242\345\214\272\345\237\237", nullptr));
        label_3->setText(QCoreApplication::translate("DataView", "\346\220\234\347\264\242\347\273\223\347\202\271  ", nullptr));
        pushButton_2->setText(QCoreApplication::translate("DataView", "\346\270\205\351\231\244", nullptr));
        pushButton->setText(QCoreApplication::translate("DataView", "\346\220\234\347\264\242", nullptr));
        GP2->setTitle(QCoreApplication::translate("DataView", "\346\225\260\346\215\256", nullptr));
        pushButton_4->setText(QCoreApplication::translate("DataView", "\346\222\244\345\233\236", nullptr));
        pushButton_3->setText(QCoreApplication::translate("DataView", "\347\241\256\345\256\232", nullptr));
        label_10->setText(QCoreApplication::translate("DataView", "\346\237\245\347\234\213\345\215\212\345\276\204", nullptr));
        label_6->setText(QCoreApplication::translate("DataView", "\346\225\260\346\215\256\346\235\245\346\272\220", nullptr));
        label_8->setText(QCoreApplication::translate("DataView", "         \345\210\260", nullptr));
        radioButton_2->setText(QCoreApplication::translate("DataView", "\346\234\215\345\212\241\345\231\250\346\225\260\346\215\256", nullptr));
        label_7->setText(QCoreApplication::translate("DataView", "         \344\273\216", nullptr));
        radioButton->setText(QCoreApplication::translate("DataView", "\346\234\254\345\234\260\346\225\260\346\215\256", nullptr));
        GP3->setTitle(QCoreApplication::translate("DataView", "\347\255\233\351\200\211", nullptr));
        checkBox_3->setText(QCoreApplication::translate("DataView", "\346\231\256\351\200\232\350\212\202\347\202\271", nullptr));
        checkBox_4->setText(QCoreApplication::translate("DataView", "\351\224\241\343\200\201\351\224\221\343\200\201\346\261\236\345\267\245\344\270\232\350\212\202\347\202\271", nullptr));
        checkBox_5->setText(QCoreApplication::translate("DataView", "\345\210\266\351\235\251\345\217\212\346\257\233\347\232\256\345\212\240\345\267\245\345\267\245\344\270\232\350\212\202\347\202\271", nullptr));
        checkBox_7->setText(QCoreApplication::translate("DataView", "\347\274\253\344\270\235\345\267\245\344\270\232\347\273\223\347\202\271", nullptr));
        radioButton_4->setText(QCoreApplication::translate("DataView", "\346\214\211\347\261\273\345\236\213\347\255\233\351\200\211", nullptr));
        pushButton_6->setText(QCoreApplication::translate("DataView", "\346\270\205\351\231\244", nullptr));
        pushButton_5->setText(QCoreApplication::translate("DataView", "\347\241\256\345\256\232", nullptr));
        checkBox_6->setText(QCoreApplication::translate("DataView", "\346\237\240\346\252\254\351\205\270\345\267\245\344\270\232\350\212\202\347\202\271", nullptr));
        radioButton_3->setText(QCoreApplication::translate("DataView", "\346\214\211\346\225\260\346\215\256\347\255\233\351\200\211", nullptr));
        checkBox->setText(QCoreApplication::translate("DataView", "\350\266\205\346\240\207\346\225\260\346\215\256", nullptr));
        checkBox_2->setText(QCoreApplication::translate("DataView", "\350\276\276\346\240\207\346\225\260\346\215\256", nullptr));
    } // retranslateUi

};

namespace Ui {
    class DataView: public Ui_DataView {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_DATAVIEW_H
