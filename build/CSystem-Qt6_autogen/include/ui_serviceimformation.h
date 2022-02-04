/********************************************************************************
** Form generated from reading UI file 'serviceimformation.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_SERVICEIMFORMATION_H
#define UI_SERVICEIMFORMATION_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_ServiceImformation
{
public:

    void setupUi(QWidget *ServiceImformation)
    {
        if (ServiceImformation->objectName().isEmpty())
            ServiceImformation->setObjectName(QString::fromUtf8("ServiceImformation"));
        ServiceImformation->resize(400, 300);

        retranslateUi(ServiceImformation);

        QMetaObject::connectSlotsByName(ServiceImformation);
    } // setupUi

    void retranslateUi(QWidget *ServiceImformation)
    {
        ServiceImformation->setWindowTitle(QCoreApplication::translate("ServiceImformation", "Form", nullptr));
    } // retranslateUi

};

namespace Ui {
    class ServiceImformation: public Ui_ServiceImformation {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_SERVICEIMFORMATION_H
