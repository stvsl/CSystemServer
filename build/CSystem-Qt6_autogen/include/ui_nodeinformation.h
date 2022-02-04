/********************************************************************************
** Form generated from reading UI file 'nodeinformation.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_NODEINFORMATION_H
#define UI_NODEINFORMATION_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QCheckBox>
#include <QtWidgets/QDateTimeEdit>
#include <QtWidgets/QGroupBox>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QSpacerItem>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_nodeinformation
{
public:
    QLabel *label;
    QWidget *layoutWidget;
    QVBoxLayout *verticalLayout_4;
    QHBoxLayout *horizontalLayout_6;
    QVBoxLayout *verticalLayout_2;
    QGroupBox *groupBox_6;
    QGroupBox *groupBox;
    QGroupBox *groupBox_2;
    QHBoxLayout *horizontalLayout_4;
    QGroupBox *groupBox_3;
    QGroupBox *groupBox_4;
    QHBoxLayout *horizontalLayout_7;
    QGroupBox *groupBox_5;
    QSpacerItem *verticalSpacer_3;
    QHBoxLayout *horizontalLayout_5;
    QVBoxLayout *verticalLayout;
    QLabel *label_2;
    QHBoxLayout *horizontalLayout;
    QLabel *label_3;
    QLineEdit *lineEdit;
    QPushButton *pushButton;
    QHBoxLayout *horizontalLayout_2;
    QLabel *label_4;
    QCheckBox *checkBox;
    QCheckBox *checkBox_2;
    QCheckBox *checkBox_3;
    QCheckBox *checkBox_4;
    QHBoxLayout *horizontalLayout_3;
    QSpacerItem *horizontalSpacer;
    QCheckBox *checkBox_5;
    QLabel *label_5;
    QDateTimeEdit *dateTimeEdit;
    QLabel *label_6;
    QDateTimeEdit *dateTimeEdit_2;
    QVBoxLayout *verticalLayout_3;
    QSpacerItem *verticalSpacer;
    QPushButton *pushButton_2;
    QPushButton *pushButton_3;
    QSpacerItem *horizontalSpacer_2;

    void setupUi(QWidget *nodeinformation)
    {
        if (nodeinformation->objectName().isEmpty())
            nodeinformation->setObjectName(QString::fromUtf8("nodeinformation"));
        nodeinformation->resize(1700, 900);
        label = new QLabel(nodeinformation);
        label->setObjectName(QString::fromUtf8("label"));
        label->setGeometry(QRect(30, 15, 91, 21));
        QFont font;
        font.setPointSize(14);
        font.setBold(true);
        label->setFont(font);
        layoutWidget = new QWidget(nodeinformation);
        layoutWidget->setObjectName(QString::fromUtf8("layoutWidget"));
        layoutWidget->setGeometry(QRect(30, 40, 1641, 857));
        verticalLayout_4 = new QVBoxLayout(layoutWidget);
        verticalLayout_4->setObjectName(QString::fromUtf8("verticalLayout_4"));
        verticalLayout_4->setContentsMargins(0, 0, 0, 0);
        horizontalLayout_6 = new QHBoxLayout();
        horizontalLayout_6->setObjectName(QString::fromUtf8("horizontalLayout_6"));
        verticalLayout_2 = new QVBoxLayout();
        verticalLayout_2->setObjectName(QString::fromUtf8("verticalLayout_2"));
        groupBox_6 = new QGroupBox(layoutWidget);
        groupBox_6->setObjectName(QString::fromUtf8("groupBox_6"));

        verticalLayout_2->addWidget(groupBox_6);


        horizontalLayout_6->addLayout(verticalLayout_2);

        groupBox = new QGroupBox(layoutWidget);
        groupBox->setObjectName(QString::fromUtf8("groupBox"));
        groupBox->setMinimumSize(QSize(900, 350));

        horizontalLayout_6->addWidget(groupBox);

        groupBox_2 = new QGroupBox(layoutWidget);
        groupBox_2->setObjectName(QString::fromUtf8("groupBox_2"));
        groupBox_2->setMinimumSize(QSize(400, 350));

        horizontalLayout_6->addWidget(groupBox_2);


        verticalLayout_4->addLayout(horizontalLayout_6);

        horizontalLayout_4 = new QHBoxLayout();
        horizontalLayout_4->setObjectName(QString::fromUtf8("horizontalLayout_4"));
        groupBox_3 = new QGroupBox(layoutWidget);
        groupBox_3->setObjectName(QString::fromUtf8("groupBox_3"));
        groupBox_3->setMinimumSize(QSize(816, 350));

        horizontalLayout_4->addWidget(groupBox_3);

        groupBox_4 = new QGroupBox(layoutWidget);
        groupBox_4->setObjectName(QString::fromUtf8("groupBox_4"));
        groupBox_4->setMinimumSize(QSize(600, 350));

        horizontalLayout_4->addWidget(groupBox_4);


        verticalLayout_4->addLayout(horizontalLayout_4);

        horizontalLayout_7 = new QHBoxLayout();
        horizontalLayout_7->setObjectName(QString::fromUtf8("horizontalLayout_7"));
        groupBox_5 = new QGroupBox(layoutWidget);
        groupBox_5->setObjectName(QString::fromUtf8("groupBox_5"));
        groupBox_5->setMinimumSize(QSize(820, 0));

        horizontalLayout_7->addWidget(groupBox_5);

        verticalSpacer_3 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        horizontalLayout_7->addItem(verticalSpacer_3);

        horizontalLayout_5 = new QHBoxLayout();
        horizontalLayout_5->setObjectName(QString::fromUtf8("horizontalLayout_5"));
        verticalLayout = new QVBoxLayout();
        verticalLayout->setObjectName(QString::fromUtf8("verticalLayout"));
        label_2 = new QLabel(layoutWidget);
        label_2->setObjectName(QString::fromUtf8("label_2"));

        verticalLayout->addWidget(label_2);

        horizontalLayout = new QHBoxLayout();
        horizontalLayout->setObjectName(QString::fromUtf8("horizontalLayout"));
        label_3 = new QLabel(layoutWidget);
        label_3->setObjectName(QString::fromUtf8("label_3"));

        horizontalLayout->addWidget(label_3);

        lineEdit = new QLineEdit(layoutWidget);
        lineEdit->setObjectName(QString::fromUtf8("lineEdit"));

        horizontalLayout->addWidget(lineEdit);

        pushButton = new QPushButton(layoutWidget);
        pushButton->setObjectName(QString::fromUtf8("pushButton"));

        horizontalLayout->addWidget(pushButton);


        verticalLayout->addLayout(horizontalLayout);

        horizontalLayout_2 = new QHBoxLayout();
        horizontalLayout_2->setObjectName(QString::fromUtf8("horizontalLayout_2"));
        label_4 = new QLabel(layoutWidget);
        label_4->setObjectName(QString::fromUtf8("label_4"));

        horizontalLayout_2->addWidget(label_4);

        checkBox = new QCheckBox(layoutWidget);
        checkBox->setObjectName(QString::fromUtf8("checkBox"));

        horizontalLayout_2->addWidget(checkBox);

        checkBox_2 = new QCheckBox(layoutWidget);
        checkBox_2->setObjectName(QString::fromUtf8("checkBox_2"));

        horizontalLayout_2->addWidget(checkBox_2);

        checkBox_3 = new QCheckBox(layoutWidget);
        checkBox_3->setObjectName(QString::fromUtf8("checkBox_3"));

        horizontalLayout_2->addWidget(checkBox_3);

        checkBox_4 = new QCheckBox(layoutWidget);
        checkBox_4->setObjectName(QString::fromUtf8("checkBox_4"));

        horizontalLayout_2->addWidget(checkBox_4);


        verticalLayout->addLayout(horizontalLayout_2);

        horizontalLayout_3 = new QHBoxLayout();
        horizontalLayout_3->setObjectName(QString::fromUtf8("horizontalLayout_3"));
        horizontalSpacer = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout_3->addItem(horizontalSpacer);

        checkBox_5 = new QCheckBox(layoutWidget);
        checkBox_5->setObjectName(QString::fromUtf8("checkBox_5"));

        horizontalLayout_3->addWidget(checkBox_5);

        label_5 = new QLabel(layoutWidget);
        label_5->setObjectName(QString::fromUtf8("label_5"));

        horizontalLayout_3->addWidget(label_5);

        dateTimeEdit = new QDateTimeEdit(layoutWidget);
        dateTimeEdit->setObjectName(QString::fromUtf8("dateTimeEdit"));

        horizontalLayout_3->addWidget(dateTimeEdit);

        label_6 = new QLabel(layoutWidget);
        label_6->setObjectName(QString::fromUtf8("label_6"));

        horizontalLayout_3->addWidget(label_6);

        dateTimeEdit_2 = new QDateTimeEdit(layoutWidget);
        dateTimeEdit_2->setObjectName(QString::fromUtf8("dateTimeEdit_2"));

        horizontalLayout_3->addWidget(dateTimeEdit_2);


        verticalLayout->addLayout(horizontalLayout_3);


        horizontalLayout_5->addLayout(verticalLayout);

        verticalLayout_3 = new QVBoxLayout();
        verticalLayout_3->setObjectName(QString::fromUtf8("verticalLayout_3"));
        verticalSpacer = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout_3->addItem(verticalSpacer);

        pushButton_2 = new QPushButton(layoutWidget);
        pushButton_2->setObjectName(QString::fromUtf8("pushButton_2"));
        pushButton_2->setMinimumSize(QSize(110, 0));
        pushButton_2->setMaximumSize(QSize(110, 16777215));

        verticalLayout_3->addWidget(pushButton_2);

        pushButton_3 = new QPushButton(layoutWidget);
        pushButton_3->setObjectName(QString::fromUtf8("pushButton_3"));
        pushButton_3->setMaximumSize(QSize(110, 16777215));

        verticalLayout_3->addWidget(pushButton_3);

        horizontalSpacer_2 = new QSpacerItem(40, 20, QSizePolicy::Minimum, QSizePolicy::Minimum);

        verticalLayout_3->addItem(horizontalSpacer_2);


        horizontalLayout_5->addLayout(verticalLayout_3);


        horizontalLayout_7->addLayout(horizontalLayout_5);


        verticalLayout_4->addLayout(horizontalLayout_7);


        retranslateUi(nodeinformation);

        QMetaObject::connectSlotsByName(nodeinformation);
    } // setupUi

    void retranslateUi(QWidget *nodeinformation)
    {
        nodeinformation->setWindowTitle(QCoreApplication::translate("nodeinformation", "Form", nullptr));
        label->setText(QCoreApplication::translate("nodeinformation", "\350\212\202\347\202\271\350\257\246\346\203\205", nullptr));
        groupBox_6->setTitle(QCoreApplication::translate("nodeinformation", "\350\212\202\347\202\271\351\200\211\346\213\251", nullptr));
        groupBox->setTitle(QCoreApplication::translate("nodeinformation", "\347\233\221\346\216\247\344\277\241\346\201\257", nullptr));
        groupBox_2->setTitle(QCoreApplication::translate("nodeinformation", "\344\275\215\347\275\256\344\277\241\346\201\257", nullptr));
        groupBox_3->setTitle(QCoreApplication::translate("nodeinformation", "\350\212\202\347\202\271\351\205\215\347\275\256\344\277\241\346\201\257", nullptr));
        groupBox_4->setTitle(QCoreApplication::translate("nodeinformation", "\350\264\237\350\264\243\346\234\272\346\236\204\345\217\212\344\272\272\345\221\230\344\277\241\346\201\257", nullptr));
        groupBox_5->setTitle(QCoreApplication::translate("nodeinformation", "\350\256\276\345\244\207\350\207\252\346\243\200\344\277\241\346\201\257", nullptr));
        label_2->setText(QCoreApplication::translate("nodeinformation", "\344\277\241\346\201\257\345\257\274\345\207\272", nullptr));
        label_3->setText(QCoreApplication::translate("nodeinformation", "\346\226\207\344\273\266\350\267\257\345\276\204\357\274\232", nullptr));
        pushButton->setText(QCoreApplication::translate("nodeinformation", "\351\200\211\346\213\251\350\267\257\345\276\204", nullptr));
        label_4->setText(QCoreApplication::translate("nodeinformation", "\345\257\274\345\207\272\345\206\205\345\256\271\357\274\232", nullptr));
        checkBox->setText(QCoreApplication::translate("nodeinformation", "\344\275\215\347\275\256\344\277\241\346\201\257", nullptr));
        checkBox_2->setText(QCoreApplication::translate("nodeinformation", "\347\233\221\346\216\247\344\277\241\346\201\257", nullptr));
        checkBox_3->setText(QCoreApplication::translate("nodeinformation", "\350\256\276\345\244\207\350\207\252\346\243\200\344\277\241\346\201\257", nullptr));
        checkBox_4->setText(QCoreApplication::translate("nodeinformation", "\350\264\237\350\264\243\344\272\272\345\221\230\344\277\241\346\201\257", nullptr));
        checkBox_5->setText(QCoreApplication::translate("nodeinformation", "\350\277\236\345\270\246\345\257\274\345\207\272\347\233\221\346\216\247\346\225\260\346\215\256", nullptr));
        label_5->setText(QCoreApplication::translate("nodeinformation", "\344\273\216", nullptr));
        label_6->setText(QCoreApplication::translate("nodeinformation", "\345\210\260", nullptr));
        pushButton_2->setText(QCoreApplication::translate("nodeinformation", "\345\257\274\345\207\272", nullptr));
        pushButton_3->setText(QCoreApplication::translate("nodeinformation", "\346\270\205\351\231\244", nullptr));
    } // retranslateUi

};

namespace Ui {
    class nodeinformation: public Ui_nodeinformation {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_NODEINFORMATION_H
