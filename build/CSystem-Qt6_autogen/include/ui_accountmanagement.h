/********************************************************************************
** Form generated from reading UI file 'accountmanagement.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_ACCOUNTMANAGEMENT_H
#define UI_ACCOUNTMANAGEMENT_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QGridLayout>
#include <QtWidgets/QGroupBox>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QHeaderView>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QRadioButton>
#include <QtWidgets/QSpacerItem>
#include <QtWidgets/QTableView>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_AccountManagement
{
public:
    QLabel *label;
    QWidget *layoutWidget;
    QGridLayout *gridLayout_5;
    QSpacerItem *horizontalSpacer_20;
    QGroupBox *groupBox;
    QGridLayout *gridLayout_3;
    QSpacerItem *horizontalSpacer_15;
    QRadioButton *radioButton;
    QSpacerItem *horizontalSpacer_16;
    QRadioButton *radioButton_2;
    QSpacerItem *horizontalSpacer_14;
    QRadioButton *radioButton_4;
    QRadioButton *radioButton_5;
    QSpacerItem *horizontalSpacer_13;
    QRadioButton *radioButton_3;
    QPushButton *pushButton_4;
    QGridLayout *gridLayout_4;
    QLabel *label_2;
    QTableView *tableView;
    QSpacerItem *horizontalSpacer_17;
    QSpacerItem *horizontalSpacer_19;
    QSpacerItem *verticalSpacer_8;
    QSpacerItem *verticalSpacer_12;
    QHBoxLayout *horizontalLayout;
    QLabel *label_3;
    QLineEdit *lineEdit_6;
    QPushButton *pushButton_16;
    QSpacerItem *verticalSpacer_9;
    QGroupBox *groupBox1;
    QGridLayout *gridLayout_2;
    QSpacerItem *horizontalSpacer_10;
    QSpacerItem *horizontalSpacer_12;
    QSpacerItem *horizontalSpacer_11;
    QLabel *label_10;
    QLabel *label_7;
    QSpacerItem *horizontalSpacer_9;
    QPushButton *pushButton_2;
    QSpacerItem *verticalSpacer_4;
    QSpacerItem *verticalSpacer_7;
    QLabel *label_8;
    QSpacerItem *horizontalSpacer_8;
    QLabel *label_9;
    QLabel *label_6;
    QSpacerItem *verticalSpacer_5;
    QSpacerItem *verticalSpacer_2;
    QSpacerItem *verticalSpacer_6;
    QPushButton *pushButton;
    QSpacerItem *horizontalSpacer_18;
    QLineEdit *lineEdit_2;
    QLineEdit *lineEdit_3;
    QLineEdit *lineEdit_4;
    QLineEdit *lineEdit_5;
    QLineEdit *lineEdit;
    QSpacerItem *verticalSpacer_10;
    QGroupBox *GP;
    QGridLayout *gridLayout;
    QPushButton *pushButton_7;
    QPushButton *pushButton_10;
    QPushButton *pushButton_3;
    QSpacerItem *horizontalSpacer_4;
    QPushButton *pushButton_13;
    QSpacerItem *horizontalSpacer_5;
    QSpacerItem *horizontalSpacer;
    QPushButton *pushButton_8;
    QPushButton *pushButton_15;
    QPushButton *pushButton_5;
    QSpacerItem *verticalSpacer_3;
    QLabel *label_12;
    QPushButton *pushButton_9;
    QSpacerItem *horizontalSpacer_7;
    QPushButton *pushButton_12;
    QSpacerItem *horizontalSpacer_2;
    QPushButton *pushButton_11;
    QSpacerItem *horizontalSpacer_6;
    QPushButton *pushButton_6;
    QSpacerItem *horizontalSpacer_3;
    QLabel *label_13;
    QSpacerItem *verticalSpacer;
    QPushButton *pushButton_14;
    QLabel *label_11;

    void setupUi(QWidget *AccountManagement)
    {
        if (AccountManagement->objectName().isEmpty())
            AccountManagement->setObjectName(QString::fromUtf8("AccountManagement"));
        AccountManagement->resize(1700, 900);
        label = new QLabel(AccountManagement);
        label->setObjectName(QString::fromUtf8("label"));
        label->setGeometry(QRect(30, 15, 91, 22));
        QFont font;
        font.setPointSize(14);
        font.setBold(true);
        font.setWeight(75);
        label->setFont(font);
        layoutWidget = new QWidget(AccountManagement);
        layoutWidget->setObjectName(QString::fromUtf8("layoutWidget"));
        layoutWidget->setGeometry(QRect(100, 60, 1523, 821));
        gridLayout_5 = new QGridLayout(layoutWidget);
        gridLayout_5->setObjectName(QString::fromUtf8("gridLayout_5"));
        gridLayout_5->setContentsMargins(0, 0, 0, 0);
        horizontalSpacer_20 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_5->addItem(horizontalSpacer_20, 4, 4, 1, 1);

        groupBox = new QGroupBox(layoutWidget);
        groupBox->setObjectName(QString::fromUtf8("groupBox"));
        gridLayout_3 = new QGridLayout(groupBox);
        gridLayout_3->setObjectName(QString::fromUtf8("gridLayout_3"));
        horizontalSpacer_15 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_3->addItem(horizontalSpacer_15, 2, 5, 1, 1);

        radioButton = new QRadioButton(groupBox);
        radioButton->setObjectName(QString::fromUtf8("radioButton"));

        gridLayout_3->addWidget(radioButton, 1, 3, 1, 3);

        horizontalSpacer_16 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_3->addItem(horizontalSpacer_16, 0, 0, 1, 1);

        radioButton_2 = new QRadioButton(groupBox);
        radioButton_2->setObjectName(QString::fromUtf8("radioButton_2"));

        gridLayout_3->addWidget(radioButton_2, 0, 1, 1, 1);

        horizontalSpacer_14 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_3->addItem(horizontalSpacer_14, 0, 2, 1, 1);

        radioButton_4 = new QRadioButton(groupBox);
        radioButton_4->setObjectName(QString::fromUtf8("radioButton_4"));

        gridLayout_3->addWidget(radioButton_4, 0, 3, 1, 3);

        radioButton_5 = new QRadioButton(groupBox);
        radioButton_5->setObjectName(QString::fromUtf8("radioButton_5"));

        gridLayout_3->addWidget(radioButton_5, 2, 1, 1, 1);

        horizontalSpacer_13 = new QSpacerItem(87, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_3->addItem(horizontalSpacer_13, 2, 3, 1, 1);

        radioButton_3 = new QRadioButton(groupBox);
        radioButton_3->setObjectName(QString::fromUtf8("radioButton_3"));

        gridLayout_3->addWidget(radioButton_3, 1, 1, 1, 1);

        pushButton_4 = new QPushButton(groupBox);
        pushButton_4->setObjectName(QString::fromUtf8("pushButton_4"));

        gridLayout_3->addWidget(pushButton_4, 2, 4, 1, 1);


        gridLayout_5->addWidget(groupBox, 2, 2, 1, 2);

        gridLayout_4 = new QGridLayout();
        gridLayout_4->setObjectName(QString::fromUtf8("gridLayout_4"));
        label_2 = new QLabel(layoutWidget);
        label_2->setObjectName(QString::fromUtf8("label_2"));

        gridLayout_4->addWidget(label_2, 0, 0, 1, 1);

        tableView = new QTableView(layoutWidget);
        tableView->setObjectName(QString::fromUtf8("tableView"));
        tableView->setMinimumSize(QSize(550, 0));

        gridLayout_4->addWidget(tableView, 1, 0, 1, 1);


        gridLayout_5->addLayout(gridLayout_4, 0, 0, 9, 1);

        horizontalSpacer_17 = new QSpacerItem(75, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_5->addItem(horizontalSpacer_17, 5, 1, 1, 1);

        horizontalSpacer_19 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_5->addItem(horizontalSpacer_19, 2, 4, 1, 1);

        verticalSpacer_8 = new QSpacerItem(17, 30, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_5->addItem(verticalSpacer_8, 1, 3, 1, 1);

        verticalSpacer_12 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_5->addItem(verticalSpacer_12, 3, 2, 1, 1);

        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setObjectName(QString::fromUtf8("horizontalLayout"));
        label_3 = new QLabel(layoutWidget);
        label_3->setObjectName(QString::fromUtf8("label_3"));

        horizontalLayout->addWidget(label_3);

        lineEdit_6 = new QLineEdit(layoutWidget);
        lineEdit_6->setObjectName(QString::fromUtf8("lineEdit_6"));

        horizontalLayout->addWidget(lineEdit_6);

        pushButton_16 = new QPushButton(layoutWidget);
        pushButton_16->setObjectName(QString::fromUtf8("pushButton_16"));

        horizontalLayout->addWidget(pushButton_16);


        gridLayout_5->addLayout(horizontalLayout, 0, 2, 1, 2);

        verticalSpacer_9 = new QSpacerItem(20, 75, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_5->addItem(verticalSpacer_9, 0, 1, 1, 1);

        groupBox1 = new QGroupBox(layoutWidget);
        groupBox1->setObjectName(QString::fromUtf8("groupBox1"));
        gridLayout_2 = new QGridLayout(groupBox1);
        gridLayout_2->setObjectName(QString::fromUtf8("gridLayout_2"));
        horizontalSpacer_10 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_2->addItem(horizontalSpacer_10, 11, 7, 1, 1);

        horizontalSpacer_12 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_2->addItem(horizontalSpacer_12, 0, 0, 1, 1);

        horizontalSpacer_11 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_2->addItem(horizontalSpacer_11, 11, 5, 1, 1);

        label_10 = new QLabel(groupBox1);
        label_10->setObjectName(QString::fromUtf8("label_10"));

        gridLayout_2->addWidget(label_10, 8, 1, 1, 1);

        label_7 = new QLabel(groupBox1);
        label_7->setObjectName(QString::fromUtf8("label_7"));

        gridLayout_2->addWidget(label_7, 2, 1, 1, 1);

        horizontalSpacer_9 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_2->addItem(horizontalSpacer_9, 11, 3, 1, 1);

        pushButton_2 = new QPushButton(groupBox1);
        pushButton_2->setObjectName(QString::fromUtf8("pushButton_2"));

        gridLayout_2->addWidget(pushButton_2, 11, 4, 1, 1);

        verticalSpacer_4 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_4, 3, 1, 1, 1);

        verticalSpacer_7 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_7, 9, 1, 1, 1);

        label_8 = new QLabel(groupBox1);
        label_8->setObjectName(QString::fromUtf8("label_8"));

        gridLayout_2->addWidget(label_8, 4, 1, 1, 1);

        horizontalSpacer_8 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_2->addItem(horizontalSpacer_8, 11, 1, 1, 1);

        label_9 = new QLabel(groupBox1);
        label_9->setObjectName(QString::fromUtf8("label_9"));

        gridLayout_2->addWidget(label_9, 6, 1, 1, 1);

        label_6 = new QLabel(groupBox1);
        label_6->setObjectName(QString::fromUtf8("label_6"));

        gridLayout_2->addWidget(label_6, 0, 1, 1, 1);

        verticalSpacer_5 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_5, 5, 1, 1, 1);

        verticalSpacer_2 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_2, 1, 1, 1, 1);

        verticalSpacer_6 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer_6, 7, 1, 1, 1);

        pushButton = new QPushButton(groupBox1);
        pushButton->setObjectName(QString::fromUtf8("pushButton"));

        gridLayout_2->addWidget(pushButton, 11, 6, 1, 1);

        horizontalSpacer_18 = new QSpacerItem(15, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_2->addItem(horizontalSpacer_18, 3, 8, 1, 1);

        lineEdit_2 = new QLineEdit(groupBox1);
        lineEdit_2->setObjectName(QString::fromUtf8("lineEdit_2"));

        gridLayout_2->addWidget(lineEdit_2, 2, 2, 1, 6);

        lineEdit_3 = new QLineEdit(groupBox1);
        lineEdit_3->setObjectName(QString::fromUtf8("lineEdit_3"));

        gridLayout_2->addWidget(lineEdit_3, 4, 2, 1, 6);

        lineEdit_4 = new QLineEdit(groupBox1);
        lineEdit_4->setObjectName(QString::fromUtf8("lineEdit_4"));

        gridLayout_2->addWidget(lineEdit_4, 6, 2, 1, 6);

        lineEdit_5 = new QLineEdit(groupBox1);
        lineEdit_5->setObjectName(QString::fromUtf8("lineEdit_5"));

        gridLayout_2->addWidget(lineEdit_5, 8, 2, 1, 6);

        lineEdit = new QLineEdit(groupBox1);
        lineEdit->setObjectName(QString::fromUtf8("lineEdit"));

        gridLayout_2->addWidget(lineEdit, 0, 2, 1, 6);


        gridLayout_5->addWidget(groupBox1, 4, 2, 2, 2);

        verticalSpacer_10 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_5->addItem(verticalSpacer_10, 6, 2, 1, 1);

        GP = new QGroupBox(layoutWidget);
        GP->setObjectName(QString::fromUtf8("GP"));
        gridLayout = new QGridLayout(GP);
        gridLayout->setObjectName(QString::fromUtf8("gridLayout"));
        pushButton_7 = new QPushButton(GP);
        pushButton_7->setObjectName(QString::fromUtf8("pushButton_7"));

        gridLayout->addWidget(pushButton_7, 5, 7, 1, 1);

        pushButton_10 = new QPushButton(GP);
        pushButton_10->setObjectName(QString::fromUtf8("pushButton_10"));

        gridLayout->addWidget(pushButton_10, 2, 11, 1, 1);

        pushButton_3 = new QPushButton(GP);
        pushButton_3->setObjectName(QString::fromUtf8("pushButton_3"));

        gridLayout->addWidget(pushButton_3, 2, 1, 1, 1);

        horizontalSpacer_4 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout->addItem(horizontalSpacer_4, 2, 8, 1, 1);

        pushButton_13 = new QPushButton(GP);
        pushButton_13->setObjectName(QString::fromUtf8("pushButton_13"));

        gridLayout->addWidget(pushButton_13, 5, 3, 1, 1);

        horizontalSpacer_5 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout->addItem(horizontalSpacer_5, 2, 10, 1, 1);

        horizontalSpacer = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout->addItem(horizontalSpacer, 2, 2, 1, 1);

        pushButton_8 = new QPushButton(GP);
        pushButton_8->setObjectName(QString::fromUtf8("pushButton_8"));

        gridLayout->addWidget(pushButton_8, 2, 7, 1, 1);

        pushButton_15 = new QPushButton(GP);
        pushButton_15->setObjectName(QString::fromUtf8("pushButton_15"));

        gridLayout->addWidget(pushButton_15, 5, 9, 1, 1);

        pushButton_5 = new QPushButton(GP);
        pushButton_5->setObjectName(QString::fromUtf8("pushButton_5"));

        gridLayout->addWidget(pushButton_5, 2, 3, 1, 1);

        verticalSpacer_3 = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer_3, 6, 7, 1, 1);

        label_12 = new QLabel(GP);
        label_12->setObjectName(QString::fromUtf8("label_12"));
        QFont font1;
        font1.setPointSize(10);
        label_12->setFont(font1);
        label_12->setStyleSheet(QString::fromUtf8("color: rgb(119, 119, 119);"));

        gridLayout->addWidget(label_12, 0, 7, 1, 1);

        pushButton_9 = new QPushButton(GP);
        pushButton_9->setObjectName(QString::fromUtf8("pushButton_9"));

        gridLayout->addWidget(pushButton_9, 2, 9, 1, 1);

        horizontalSpacer_7 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout->addItem(horizontalSpacer_7, 2, 12, 1, 1);

        pushButton_12 = new QPushButton(GP);
        pushButton_12->setObjectName(QString::fromUtf8("pushButton_12"));

        gridLayout->addWidget(pushButton_12, 5, 1, 1, 1);

        horizontalSpacer_2 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout->addItem(horizontalSpacer_2, 2, 4, 1, 1);

        pushButton_11 = new QPushButton(GP);
        pushButton_11->setObjectName(QString::fromUtf8("pushButton_11"));

        gridLayout->addWidget(pushButton_11, 5, 5, 1, 1);

        horizontalSpacer_6 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout->addItem(horizontalSpacer_6, 2, 0, 1, 1);

        pushButton_6 = new QPushButton(GP);
        pushButton_6->setObjectName(QString::fromUtf8("pushButton_6"));

        gridLayout->addWidget(pushButton_6, 2, 5, 1, 1);

        horizontalSpacer_3 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout->addItem(horizontalSpacer_3, 2, 6, 1, 1);

        label_13 = new QLabel(GP);
        label_13->setObjectName(QString::fromUtf8("label_13"));
        label_13->setFont(font1);
        label_13->setStyleSheet(QString::fromUtf8("color: rgb(119, 119, 119);"));

        gridLayout->addWidget(label_13, 4, 7, 1, 1);

        verticalSpacer = new QSpacerItem(20, 20, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer, 3, 7, 1, 1);

        pushButton_14 = new QPushButton(GP);
        pushButton_14->setObjectName(QString::fromUtf8("pushButton_14"));

        gridLayout->addWidget(pushButton_14, 5, 11, 1, 1);

        label_11 = new QLabel(GP);
        label_11->setObjectName(QString::fromUtf8("label_11"));
        label_11->setFont(font1);
        label_11->setStyleSheet(QString::fromUtf8("color: rgb(119, 119, 119);"));

        gridLayout->addWidget(label_11, 0, 1, 1, 1);


        gridLayout_5->addWidget(GP, 7, 2, 1, 3);


        retranslateUi(AccountManagement);

        QMetaObject::connectSlotsByName(AccountManagement);
    } // setupUi

    void retranslateUi(QWidget *AccountManagement)
    {
        AccountManagement->setWindowTitle(QCoreApplication::translate("AccountManagement", "Form", nullptr));
        label->setText(QCoreApplication::translate("AccountManagement", "\350\264\246\346\210\267\347\256\241\347\220\206", nullptr));
        groupBox->setTitle(QCoreApplication::translate("AccountManagement", "\350\264\246\346\210\267\347\261\273\345\236\213", nullptr));
        radioButton->setText(QCoreApplication::translate("AccountManagement", "\346\231\256\351\200\232\347\256\241\347\220\206\345\221\230\350\264\246\346\210\267", nullptr));
        radioButton_2->setText(QCoreApplication::translate("AccountManagement", "\346\231\256\351\200\232\350\264\246\346\210\267", nullptr));
        radioButton_4->setText(QCoreApplication::translate("AccountManagement", "\345\256\211\350\243\205\347\273\264\346\212\244\344\272\272\345\221\230\350\264\246\346\210\267", nullptr));
        radioButton_5->setText(QCoreApplication::translate("AccountManagement", "\346\234\200\351\253\230\347\256\241\347\220\206\345\221\230\350\264\246\346\210\267", nullptr));
        radioButton_3->setText(QCoreApplication::translate("AccountManagement", "\346\225\260\346\215\256\345\275\225\345\205\245\346\234\272\346\236\204\350\264\246\346\210\267", nullptr));
        pushButton_4->setText(QCoreApplication::translate("AccountManagement", "\344\277\256\346\224\271/\350\256\276\347\275\256", nullptr));
        label_2->setText(QCoreApplication::translate("AccountManagement", "\350\264\246\346\210\267\345\210\227\350\241\250", nullptr));
        label_3->setText(QCoreApplication::translate("AccountManagement", "\346\220\234\347\264\242\350\264\246\346\210\267", nullptr));
        pushButton_16->setText(QCoreApplication::translate("AccountManagement", "\346\237\245\350\257\242", nullptr));
        groupBox1->setTitle(QCoreApplication::translate("AccountManagement", "\350\264\246\346\210\267\344\277\241\346\201\257", nullptr));
        label_10->setText(QCoreApplication::translate("AccountManagement", "\346\211\200\345\261\236\344\274\201\344\270\232/\346\234\272\346\236\204 ", nullptr));
        label_7->setText(QCoreApplication::translate("AccountManagement", "\350\272\253\344\273\275\350\257\201\345\217\267", nullptr));
        pushButton_2->setText(QCoreApplication::translate("AccountManagement", "\346\222\244\351\224\200", nullptr));
        label_8->setText(QCoreApplication::translate("AccountManagement", "\345\256\266\345\272\255\344\275\217\345\235\200", nullptr));
        label_9->setText(QCoreApplication::translate("AccountManagement", "\350\201\224\347\263\273\347\224\265\350\257\235", nullptr));
        label_6->setText(QCoreApplication::translate("AccountManagement", "\347\224\250\346\210\267\345\247\223\345\220\215", nullptr));
        pushButton->setText(QCoreApplication::translate("AccountManagement", "\344\277\256\346\224\271/\350\256\276\347\275\256", nullptr));
        GP->setTitle(QCoreApplication::translate("AccountManagement", "\346\223\215\344\275\234\344\270\255\345\277\203", nullptr));
        pushButton_7->setText(QCoreApplication::translate("AccountManagement", "\344\274\201\344\270\232/\346\234\272\346\236\204\345\210\227\350\241\250", nullptr));
        pushButton_10->setText(QCoreApplication::translate("AccountManagement", "\346\237\245\347\234\213\346\223\215\344\275\234\346\227\245\345\277\227", nullptr));
        pushButton_3->setText(QCoreApplication::translate("AccountManagement", "\346\267\273\345\212\240\350\264\246\346\210\267", nullptr));
        pushButton_13->setText(QCoreApplication::translate("AccountManagement", "\350\247\243\345\206\273\350\264\246\346\210\267", nullptr));
        pushButton_8->setText(QCoreApplication::translate("AccountManagement", "\346\237\245\347\234\213RSA\345\257\206\351\222\245\346\214\207\347\272\271", nullptr));
        pushButton_15->setText(QCoreApplication::translate("AccountManagement", "\344\274\201\344\270\232/\346\234\272\346\236\204\346\237\245\350\257\242", nullptr));
        pushButton_5->setText(QCoreApplication::translate("AccountManagement", "\346\263\250\351\224\200\350\264\246\346\210\267", nullptr));
        label_12->setText(QCoreApplication::translate("AccountManagement", "\347\275\256\344\277\241\346\223\215\344\275\234", nullptr));
        pushButton_9->setText(QCoreApplication::translate("AccountManagement", "\351\207\215\345\210\206\351\205\215RSA\345\257\206\351\222\245", nullptr));
        pushButton_12->setText(QCoreApplication::translate("AccountManagement", "\345\206\273\347\273\223\350\264\246\346\210\267", nullptr));
        pushButton_11->setText(QCoreApplication::translate("AccountManagement", "\351\207\215\347\275\256\350\264\246\346\210\267", nullptr));
        pushButton_6->setText(QCoreApplication::translate("AccountManagement", "\345\257\206\347\240\201\351\207\215\347\275\256", nullptr));
        label_13->setText(QCoreApplication::translate("AccountManagement", "\346\234\272\346\236\204\346\223\215\344\275\234", nullptr));
        pushButton_14->setText(QCoreApplication::translate("AccountManagement", "\346\267\273\345\212\240\344\274\201\344\270\232/\346\234\272\346\236\204", nullptr));
        label_11->setText(QCoreApplication::translate("AccountManagement", "\350\264\246\346\210\267\346\223\215\344\275\234", nullptr));
    } // retranslateUi

};

namespace Ui {
    class AccountManagement: public Ui_AccountManagement {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_ACCOUNTMANAGEMENT_H
