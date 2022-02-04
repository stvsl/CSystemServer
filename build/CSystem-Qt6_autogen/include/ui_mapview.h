/********************************************************************************
** Form generated from reading UI file 'mapview.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_MAPVIEW_H
#define UI_MAPVIEW_H

#include <QtCore/QVariant>
#include <QtWebEngineWidgets/QWebEngineView>
#include <QtWidgets/QApplication>
#include <QtWidgets/QGridLayout>
#include <QtWidgets/QHeaderView>
#include <QtWidgets/QLabel>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QListView>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QSpacerItem>
#include <QtWidgets/QSplitter>
#include <QtWidgets/QTableView>
#include <QtWidgets/QVBoxLayout>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_MapView
{
public:
    QLabel *label;
    QWebEngineView *webView;
    QSplitter *splitter_5;
    QSplitter *splitter_2;
    QWidget *layoutWidget;
    QGridLayout *gridLayout_2;
    QGridLayout *gridLayout;
    QLabel *label_3;
    QLabel *label_2;
    QLineEdit *nodeSearcher;
    QLineEdit *mapSearcher;
    QLabel *label_4;
    QWidget *layoutWidget1;
    QVBoxLayout *verticalLayout;
    QSpacerItem *verticalSpacer;
    QPushButton *PositionReset;
    QSplitter *splitter_3;
    QSplitter *splitter;
    QLabel *label_5;
    QListView *NodeList;
    QSplitter *splitter_4;
    QLabel *label_7;
    QTableView *tableView;

    void setupUi(QWidget *MapView)
    {
        if (MapView->objectName().isEmpty())
            MapView->setObjectName(QString::fromUtf8("MapView"));
        MapView->resize(1700, 900);
        label = new QLabel(MapView);
        label->setObjectName(QString::fromUtf8("label"));
        label->setGeometry(QRect(30, 15, 81, 21));
        QFont font;
        font.setPointSize(14);
        font.setBold(true);
        label->setFont(font);
        webView = new QWebEngineView(MapView);
        webView->setObjectName(QString::fromUtf8("webView"));
        webView->setGeometry(QRect(30, 50, 1300, 850));
        webView->setUrl(QUrl(QString::fromUtf8("about:blank")));
        splitter_5 = new QSplitter(MapView);
        splitter_5->setObjectName(QString::fromUtf8("splitter_5"));
        splitter_5->setGeometry(QRect(1350, 51, 321, 861));
        splitter_5->setOrientation(Qt::Vertical);
        splitter_2 = new QSplitter(splitter_5);
        splitter_2->setObjectName(QString::fromUtf8("splitter_2"));
        splitter_2->setOrientation(Qt::Horizontal);
        splitter_2->setOpaqueResize(false);
        splitter_2->setChildrenCollapsible(true);
        layoutWidget = new QWidget(splitter_2);
        layoutWidget->setObjectName(QString::fromUtf8("layoutWidget"));
        gridLayout_2 = new QGridLayout(layoutWidget);
        gridLayout_2->setObjectName(QString::fromUtf8("gridLayout_2"));
        gridLayout_2->setContentsMargins(0, 0, 0, 0);
        gridLayout = new QGridLayout();
        gridLayout->setObjectName(QString::fromUtf8("gridLayout"));
        label_3 = new QLabel(layoutWidget);
        label_3->setObjectName(QString::fromUtf8("label_3"));

        gridLayout->addWidget(label_3, 1, 0, 1, 1);

        label_2 = new QLabel(layoutWidget);
        label_2->setObjectName(QString::fromUtf8("label_2"));

        gridLayout->addWidget(label_2, 0, 0, 1, 1);

        nodeSearcher = new QLineEdit(layoutWidget);
        nodeSearcher->setObjectName(QString::fromUtf8("nodeSearcher"));
        nodeSearcher->setMinimumSize(QSize(149, 0));

        gridLayout->addWidget(nodeSearcher, 1, 1, 1, 1);

        mapSearcher = new QLineEdit(layoutWidget);
        mapSearcher->setObjectName(QString::fromUtf8("mapSearcher"));

        gridLayout->addWidget(mapSearcher, 0, 1, 1, 1);


        gridLayout_2->addLayout(gridLayout, 1, 0, 1, 1);

        label_4 = new QLabel(layoutWidget);
        label_4->setObjectName(QString::fromUtf8("label_4"));
        QFont font1;
        font1.setPointSize(12);
        label_4->setFont(font1);

        gridLayout_2->addWidget(label_4, 0, 0, 1, 1);

        splitter_2->addWidget(layoutWidget);
        layoutWidget1 = new QWidget(splitter_2);
        layoutWidget1->setObjectName(QString::fromUtf8("layoutWidget1"));
        verticalLayout = new QVBoxLayout(layoutWidget1);
        verticalLayout->setObjectName(QString::fromUtf8("verticalLayout"));
        verticalLayout->setContentsMargins(0, 0, 0, 0);
        verticalSpacer = new QSpacerItem(20, 23, QSizePolicy::Minimum, QSizePolicy::Expanding);

        verticalLayout->addItem(verticalSpacer);

        PositionReset = new QPushButton(layoutWidget1);
        PositionReset->setObjectName(QString::fromUtf8("PositionReset"));
        PositionReset->setMinimumSize(QSize(0, 55));
        PositionReset->setMaximumSize(QSize(16777215, 55));

        verticalLayout->addWidget(PositionReset);

        splitter_2->addWidget(layoutWidget1);
        splitter_5->addWidget(splitter_2);
        splitter_3 = new QSplitter(splitter_5);
        splitter_3->setObjectName(QString::fromUtf8("splitter_3"));
        splitter_3->setOrientation(Qt::Vertical);
        splitter = new QSplitter(splitter_3);
        splitter->setObjectName(QString::fromUtf8("splitter"));
        splitter->setOrientation(Qt::Vertical);
        splitter_3->addWidget(splitter);
        label_5 = new QLabel(splitter_3);
        label_5->setObjectName(QString::fromUtf8("label_5"));
        label_5->setFont(font1);
        splitter_3->addWidget(label_5);
        NodeList = new QListView(splitter_3);
        NodeList->setObjectName(QString::fromUtf8("NodeList"));
        NodeList->setMinimumSize(QSize(0, 300));
        splitter_3->addWidget(NodeList);
        splitter_5->addWidget(splitter_3);
        splitter_4 = new QSplitter(splitter_5);
        splitter_4->setObjectName(QString::fromUtf8("splitter_4"));
        splitter_4->setOrientation(Qt::Vertical);
        splitter_4->setOpaqueResize(true);
        splitter_4->setHandleWidth(7);
        label_7 = new QLabel(splitter_4);
        label_7->setObjectName(QString::fromUtf8("label_7"));
        label_7->setFont(font1);
        splitter_4->addWidget(label_7);
        tableView = new QTableView(splitter_4);
        tableView->setObjectName(QString::fromUtf8("tableView"));
        tableView->setMinimumSize(QSize(0, 400));
        splitter_4->addWidget(tableView);
        splitter_5->addWidget(splitter_4);

        retranslateUi(MapView);

        QMetaObject::connectSlotsByName(MapView);
    } // setupUi

    void retranslateUi(QWidget *MapView)
    {
        MapView->setWindowTitle(QCoreApplication::translate("MapView", "Form", nullptr));
        label->setText(QCoreApplication::translate("MapView", "\345\205\250\345\261\200\350\247\206\345\233\276", nullptr));
        label_3->setText(QCoreApplication::translate("MapView", "\346\237\245\346\211\276\347\273\223\347\202\271 :  ", nullptr));
        label_2->setText(QCoreApplication::translate("MapView", "\346\237\245\346\211\276\344\275\215\347\275\256 :  ", nullptr));
#if QT_CONFIG(tooltip)
        nodeSearcher->setToolTip(QCoreApplication::translate("MapView", "\350\276\223\345\205\245\350\212\202\347\202\271\345\217\267/\345\234\260\345\220\215\350\277\233\350\241\214\346\220\234\347\264\242", nullptr));
#endif // QT_CONFIG(tooltip)
#if QT_CONFIG(tooltip)
        mapSearcher->setToolTip(QCoreApplication::translate("MapView", "\350\276\223\345\205\245\350\212\202\347\202\271\345\217\267/\345\234\260\345\220\215\350\277\233\350\241\214\346\220\234\347\264\242", nullptr));
#endif // QT_CONFIG(tooltip)
        label_4->setText(QCoreApplication::translate("MapView", "\346\237\245\346\211\276 ", nullptr));
        PositionReset->setText(QCoreApplication::translate("MapView", "\344\275\215\347\275\256\351\207\215\347\275\256", nullptr));
        label_5->setText(QCoreApplication::translate("MapView", "\350\212\202\347\202\271\345\210\227\350\241\250", nullptr));
        label_7->setText(QCoreApplication::translate("MapView", "\350\212\202\347\202\271\350\257\246\346\203\205", nullptr));
    } // retranslateUi

};

namespace Ui {
    class MapView: public Ui_MapView {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_MAPVIEW_H
