/********************************************************************************
** Form generated from reading UI file 'servicecontrol.ui'
**
** Created by: Qt User Interface Compiler version 5.15.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_SERVICECONTROL_H
#define UI_SERVICECONTROL_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_ServiceControl
{
public:

    void setupUi(QWidget *ServiceControl)
    {
        if (ServiceControl->objectName().isEmpty())
            ServiceControl->setObjectName(QString::fromUtf8("ServiceControl"));
        ServiceControl->resize(400, 300);

        retranslateUi(ServiceControl);

        QMetaObject::connectSlotsByName(ServiceControl);
    } // setupUi

    void retranslateUi(QWidget *ServiceControl)
    {
        ServiceControl->setWindowTitle(QCoreApplication::translate("ServiceControl", "Form", nullptr));
    } // retranslateUi

};

namespace Ui {
    class ServiceControl: public Ui_ServiceControl {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_SERVICECONTROL_H
