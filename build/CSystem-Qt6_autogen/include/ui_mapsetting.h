/********************************************************************************
** Form generated from reading UI file 'mapsetting.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_MAPSETTING_H
#define UI_MAPSETTING_H

#include <QtCore/QVariant>
#include <QtWebEngineWidgets/QWebEngineView>
#include <QtWidgets/QApplication>
#include <QtWidgets/QCheckBox>
#include <QtWidgets/QGridLayout>
#include <QtWidgets/QGroupBox>
#include <QtWidgets/QHBoxLayout>
#include <QtWidgets/QLabel>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QSpacerItem>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_MapSetting
{
public:
    QLabel *label;
    QWidget *layoutWidget;
    QHBoxLayout *horizontalLayout;
    QVBoxLayout *verticalLayout;
    QGridLayout *gridLayout_3;
    QSpacerItem *verticalSpacer_4;
    QCheckBox *checkBox_4;
    QLabel *label_3;
    QCheckBox *checkBox_5;
    QSpacerItem *verticalSpacer_14;
    QSpacerItem *verticalSpacer_7;
    QGridLayout *gridLayout_4;
    QLabel *label_2;
    QSpacerItem *verticalSpacer_6;
    QCheckBox *checkBox;
    QCheckBox *checkBox_2;
    QSpacerItem *verticalSpacer_9;
    QGridLayout *gridLayout_6;
    QLabel *label_7;
    QCheckBox *checkBox_11;
    QSpacerItem *verticalSpacer_5;
    QCheckBox *checkBox_10;
    QCheckBox *checkBox_9;
    QCheckBox *checkBox_7;
    QSpacerItem *verticalSpacer_13;
    QGridLayout *gridLayout_2;
    QSpacerItem *verticalSpacer;
    QCheckBox *checkBox_8;
    QLabel *label_6;
    QCheckBox *checkBox_12;
    QSpacerItem *verticalSpacer_10;
    QGridLayout *gridLayout_5;
    QLabel *label_4;
    QSpacerItem *verticalSpacer_3;
    QCheckBox *checkBox_3;
    QSpacerItem *verticalSpacer_8;
    QGridLayout *gridLayout;
    QSpacerItem *verticalSpacer_12;
    QLabel *label_5;
    QSpacerItem *verticalSpacer_2;
    QCheckBox *checkBox_6;
    QSpacerItem *horizontalSpacer_2;
    QVBoxLayout *verticalLayout_2;
    QGridLayout *gridLayout_8;
    QWebEngineView *webView;
    QSpacerItem *horizontalSpacer;
    QLabel *label_9;
    QSpacerItem *verticalSpacer_15;
    QGroupBox *groupBox;
    QLabel *label_8;
    QSpacerItem *verticalSpacer_11;
    QGridLayout *gridLayout_7;
    QSpacerItem *horizontalSpacer_3;
    QPushButton *pushButton_3;
    QPushButton *pushButton;
    QPushButton *pushButton_2;

    void setupUi(QWidget *MapSetting)
    {
        if (MapSetting->objectName().isEmpty())
            MapSetting->setObjectName(QString::fromUtf8("MapSetting"));
        MapSetting->resize(1700, 900);
        label = new QLabel(MapSetting);
        label->setObjectName(QString::fromUtf8("label"));
        label->setGeometry(QRect(30, 15, 81, 21));
        QFont font;
        font.setPointSize(14);
        font.setBold(true);
        label->setFont(font);
        layoutWidget = new QWidget(MapSetting);
        layoutWidget->setObjectName(QString::fromUtf8("layoutWidget"));
        layoutWidget->setGeometry(QRect(150, 60, 1521, 832));
        horizontalLayout = new QHBoxLayout(layoutWidget);
        horizontalLayout->setObjectName(QString::fromUtf8("horizontalLayout"));
        horizontalLayout->setContentsMargins(0, 0, 0, 0);
        verticalLayout = new QVBoxLayout();
        verticalLayout->setObjectName(QString::fromUtf8("verticalLayout"));
        gridLayout_3 = new QGridLayout();
        gridLayout_3->setObjectName(QString::fromUtf8("gridLayout_3"));
        verticalSpacer_4 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_3->addItem(verticalSpacer_4, 2, 0, 2, 1);

        checkBox_4 = new QCheckBox(layoutWidget);
        checkBox_4->setObjectName(QString::fromUtf8("checkBox_4"));

        gridLayout_3->addWidget(checkBox_4, 2, 1, 1, 1);

        label_3 = new QLabel(layoutWidget);
        label_3->setObjectName(QString::fromUtf8("label_3"));
        QFont font1;
        font1.setPointSize(12);
        label_3->setFont(font1);

        gridLayout_3->addWidget(label_3, 1, 0, 1, 2);

        checkBox_5 = new QCheckBox(layoutWidget);
        checkBox_5->setObjectName(QString::fromUtf8("checkBox_5"));

        gridLayout_3->addWidget(checkBox_5, 3, 1, 1, 1);

        verticalSpacer_14 = new QSpacerItem(20, 22, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_3->addItem(verticalSpacer_14, 0, 0, 1, 1);


        verticalLayout->addLayout(gridLayout_3);

        verticalSpacer_7 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer_7);

        gridLayout_4 = new QGridLayout();
        gridLayout_4->setObjectName(QString::fromUtf8("gridLayout_4"));
        label_2 = new QLabel(layoutWidget);
        label_2->setObjectName(QString::fromUtf8("label_2"));
        label_2->setFont(font1);

        gridLayout_4->addWidget(label_2, 0, 0, 1, 2);

        verticalSpacer_6 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_4->addItem(verticalSpacer_6, 1, 0, 2, 1);

        checkBox = new QCheckBox(layoutWidget);
        checkBox->setObjectName(QString::fromUtf8("checkBox"));

        gridLayout_4->addWidget(checkBox, 1, 1, 1, 1);

        checkBox_2 = new QCheckBox(layoutWidget);
        checkBox_2->setObjectName(QString::fromUtf8("checkBox_2"));

        gridLayout_4->addWidget(checkBox_2, 2, 1, 1, 1);


        verticalLayout->addLayout(gridLayout_4);

        verticalSpacer_9 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer_9);

        gridLayout_6 = new QGridLayout();
        gridLayout_6->setObjectName(QString::fromUtf8("gridLayout_6"));
        label_7 = new QLabel(layoutWidget);
        label_7->setObjectName(QString::fromUtf8("label_7"));
        label_7->setFont(font1);

        gridLayout_6->addWidget(label_7, 0, 0, 1, 2);

        checkBox_11 = new QCheckBox(layoutWidget);
        checkBox_11->setObjectName(QString::fromUtf8("checkBox_11"));

        gridLayout_6->addWidget(checkBox_11, 4, 1, 1, 1);

        verticalSpacer_5 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_6->addItem(verticalSpacer_5, 1, 0, 4, 1);

        checkBox_10 = new QCheckBox(layoutWidget);
        checkBox_10->setObjectName(QString::fromUtf8("checkBox_10"));

        gridLayout_6->addWidget(checkBox_10, 1, 1, 1, 1);

        checkBox_9 = new QCheckBox(layoutWidget);
        checkBox_9->setObjectName(QString::fromUtf8("checkBox_9"));

        gridLayout_6->addWidget(checkBox_9, 3, 1, 1, 1);

        checkBox_7 = new QCheckBox(layoutWidget);
        checkBox_7->setObjectName(QString::fromUtf8("checkBox_7"));

        gridLayout_6->addWidget(checkBox_7, 2, 1, 1, 1);


        verticalLayout->addLayout(gridLayout_6);

        verticalSpacer_13 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer_13);

        gridLayout_2 = new QGridLayout();
        gridLayout_2->setObjectName(QString::fromUtf8("gridLayout_2"));
        verticalSpacer = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_2->addItem(verticalSpacer, 1, 0, 2, 1);

        checkBox_8 = new QCheckBox(layoutWidget);
        checkBox_8->setObjectName(QString::fromUtf8("checkBox_8"));
        checkBox_8->setChecked(false);

        gridLayout_2->addWidget(checkBox_8, 1, 1, 1, 1);

        label_6 = new QLabel(layoutWidget);
        label_6->setObjectName(QString::fromUtf8("label_6"));
        label_6->setFont(font1);

        gridLayout_2->addWidget(label_6, 0, 0, 1, 2);

        checkBox_12 = new QCheckBox(layoutWidget);
        checkBox_12->setObjectName(QString::fromUtf8("checkBox_12"));

        gridLayout_2->addWidget(checkBox_12, 2, 1, 1, 1);


        verticalLayout->addLayout(gridLayout_2);

        verticalSpacer_10 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer_10);

        gridLayout_5 = new QGridLayout();
        gridLayout_5->setObjectName(QString::fromUtf8("gridLayout_5"));
        label_4 = new QLabel(layoutWidget);
        label_4->setObjectName(QString::fromUtf8("label_4"));
        label_4->setFont(font1);

        gridLayout_5->addWidget(label_4, 0, 0, 1, 2);

        verticalSpacer_3 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_5->addItem(verticalSpacer_3, 1, 0, 2, 1);

        checkBox_3 = new QCheckBox(layoutWidget);
        checkBox_3->setObjectName(QString::fromUtf8("checkBox_3"));

        gridLayout_5->addWidget(checkBox_3, 1, 1, 1, 1);


        verticalLayout->addLayout(gridLayout_5);

        verticalSpacer_8 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer_8);

        gridLayout = new QGridLayout();
        gridLayout->setObjectName(QString::fromUtf8("gridLayout"));
        verticalSpacer_12 = new QSpacerItem(20, 108, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer_12, 2, 1, 1, 1);

        label_5 = new QLabel(layoutWidget);
        label_5->setObjectName(QString::fromUtf8("label_5"));
        label_5->setFont(font1);

        gridLayout->addWidget(label_5, 0, 0, 1, 2);

        verticalSpacer_2 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout->addItem(verticalSpacer_2, 1, 0, 1, 1);

        checkBox_6 = new QCheckBox(layoutWidget);
        checkBox_6->setObjectName(QString::fromUtf8("checkBox_6"));

        gridLayout->addWidget(checkBox_6, 1, 1, 1, 1);


        verticalLayout->addLayout(gridLayout);


        horizontalLayout->addLayout(verticalLayout);

        horizontalSpacer_2 = new QSpacerItem(19, 19, QSizePolicy::Expanding, QSizePolicy::Minimum);

        horizontalLayout->addItem(horizontalSpacer_2);

        verticalLayout_2 = new QVBoxLayout();
        verticalLayout_2->setObjectName(QString::fromUtf8("verticalLayout_2"));
        gridLayout_8 = new QGridLayout();
        gridLayout_8->setObjectName(QString::fromUtf8("gridLayout_8"));
        webView = new QWebEngineView(layoutWidget);
        webView->setObjectName(QString::fromUtf8("webView"));
        webView->setMinimumSize(QSize(1200, 620));
        webView->setUrl(QUrl(QString::fromUtf8("about:blank")));

        gridLayout_8->addWidget(webView, 1, 0, 1, 2);

        horizontalSpacer = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_8->addItem(horizontalSpacer, 0, 1, 1, 1);

        label_9 = new QLabel(layoutWidget);
        label_9->setObjectName(QString::fromUtf8("label_9"));
        label_9->setFont(font1);

        gridLayout_8->addWidget(label_9, 0, 0, 1, 1);

        verticalSpacer_15 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        gridLayout_8->addItem(verticalSpacer_15, 2, 0, 1, 1);


        verticalLayout_2->addLayout(gridLayout_8);

        groupBox = new QGroupBox(layoutWidget);
        groupBox->setObjectName(QString::fromUtf8("groupBox"));
        groupBox->setMinimumSize(QSize(0, 120));
        label_8 = new QLabel(groupBox);
        label_8->setObjectName(QString::fromUtf8("label_8"));
        label_8->setGeometry(QRect(0, 20, 961, 91));

        verticalLayout_2->addWidget(groupBox);

        verticalSpacer_11 = new QSpacerItem(20, 40, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout_2->addItem(verticalSpacer_11);

        gridLayout_7 = new QGridLayout();
        gridLayout_7->setObjectName(QString::fromUtf8("gridLayout_7"));
        gridLayout_7->setSizeConstraint(QLayout::SetMinimumSize);
        horizontalSpacer_3 = new QSpacerItem(40, 20, QSizePolicy::Expanding, QSizePolicy::Minimum);

        gridLayout_7->addItem(horizontalSpacer_3, 0, 0, 1, 1);

        pushButton_3 = new QPushButton(layoutWidget);
        pushButton_3->setObjectName(QString::fromUtf8("pushButton_3"));

        gridLayout_7->addWidget(pushButton_3, 0, 2, 1, 1);

        pushButton = new QPushButton(layoutWidget);
        pushButton->setObjectName(QString::fromUtf8("pushButton"));

        gridLayout_7->addWidget(pushButton, 0, 1, 1, 1);

        pushButton_2 = new QPushButton(layoutWidget);
        pushButton_2->setObjectName(QString::fromUtf8("pushButton_2"));

        gridLayout_7->addWidget(pushButton_2, 0, 3, 1, 1);


        verticalLayout_2->addLayout(gridLayout_7);


        horizontalLayout->addLayout(verticalLayout_2);


        retranslateUi(MapSetting);

        QMetaObject::connectSlotsByName(MapSetting);
    } // setupUi

    void retranslateUi(QWidget *MapSetting)
    {
        MapSetting->setWindowTitle(QCoreApplication::translate("MapSetting", "Form", nullptr));
        label->setText(QCoreApplication::translate("MapSetting", "\345\234\260\345\233\276\350\256\276\347\275\256", nullptr));
        checkBox_4->setText(QCoreApplication::translate("MapSetting", "3D\346\216\247\344\273\266\346\224\257\346\214\201", nullptr));
        label_3->setText(QCoreApplication::translate("MapSetting", "\345\234\260\345\233\276\346\216\247\344\273\266\350\256\276\347\275\256", nullptr));
        checkBox_5->setText(QCoreApplication::translate("MapSetting", "\346\257\224\344\276\213\345\260\272\346\216\247\344\273\266\346\224\257\346\214\201", nullptr));
        label_2->setText(QCoreApplication::translate("MapSetting", "\345\234\260\345\233\276\345\234\260\346\240\207\350\256\276\347\275\256", nullptr));
        checkBox->setText(QCoreApplication::translate("MapSetting", "\346\230\276\347\244\272\345\234\260\346\240\207\346\240\207\350\256\260", nullptr));
        checkBox_2->setText(QCoreApplication::translate("MapSetting", "\346\230\276\347\244\272\345\234\260\346\240\207\346\226\207\345\255\227", nullptr));
        label_7->setText(QCoreApplication::translate("MapSetting", "\345\234\260\345\233\276\346\223\215\344\275\234", nullptr));
        checkBox_11->setText(QCoreApplication::translate("MapSetting", "\345\220\257\347\224\250\346\203\257\346\200\247\346\213\226\346\213\275", nullptr));
        checkBox_10->setText(QCoreApplication::translate("MapSetting", "\351\224\256\347\233\230\346\216\247\345\210\266", nullptr));
        checkBox_9->setText(QCoreApplication::translate("MapSetting", "\345\220\257\347\224\250\346\213\226\346\213\275", nullptr));
        checkBox_7->setText(QCoreApplication::translate("MapSetting", "\351\274\240\346\240\207\346\273\232\350\275\256\347\274\251\346\224\276", nullptr));
        checkBox_8->setText(QCoreApplication::translate("MapSetting", "\347\247\273\351\231\244\345\234\260\345\233\276\345\212\250\347\224\273", nullptr));
        label_6->setText(QCoreApplication::translate("MapSetting", "\345\234\260\345\233\276\346\200\247\350\203\275\344\274\230\345\214\226", nullptr));
        checkBox_12->setText(QCoreApplication::translate("MapSetting", "\345\206\205\347\275\256\346\216\247\345\210\266\345\231\250", nullptr));
        label_4->setText(QCoreApplication::translate("MapSetting", "\345\234\260\345\233\276\346\250\241\345\274\217\357\274\210\351\207\215\345\220\257\347\224\237\346\225\210\357\274\211", nullptr));
        checkBox_3->setText(QCoreApplication::translate("MapSetting", "\345\215\253\346\230\237\345\234\260\347\220\203\346\250\241\345\274\217", nullptr));
        label_5->setText(QCoreApplication::translate("MapSetting", "\345\234\260\345\233\276\345\256\232\344\275\215\357\274\210\351\207\215\345\220\257\347\224\237\346\225\210\357\274\211", nullptr));
        checkBox_6->setText(QCoreApplication::translate("MapSetting", "\351\224\232\347\202\271\350\207\252\345\212\250\345\256\232\344\275\215", nullptr));
        label_9->setText(QCoreApplication::translate("MapSetting", "\346\225\210\346\236\234\351\242\204\350\247\210", nullptr));
        groupBox->setTitle(QCoreApplication::translate("MapSetting", "\345\212\237\350\203\275\346\217\217\350\277\260", nullptr));
        label_8->setText(QCoreApplication::translate("MapSetting", "\350\277\231\346\230\257\346\265\213\350\257\225\347\232\204\345\212\237\350\203\275\346\217\217\350\277\260", nullptr));
        pushButton_3->setText(QCoreApplication::translate("MapSetting", "\345\217\226\346\266\210", nullptr));
        pushButton->setText(QCoreApplication::translate("MapSetting", "\351\207\215\347\275\256", nullptr));
        pushButton_2->setText(QCoreApplication::translate("MapSetting", "\347\241\256\350\256\244", nullptr));
    } // retranslateUi

};

namespace Ui {
    class MapSetting: public Ui_MapSetting {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_MAPSETTING_H
