#ifndef DBUTILS_H
#define DBUTILS_H

#include <QSqlDatabase>
#include <QSqlError>
#include <QSqlQuery>
#include "configManager/configmanager.h"

class DButils
{

public:
    DButils();

private:
    //数据库连接
    QSqlDatabase db;
    //
};

#endif // DBUTILS_H