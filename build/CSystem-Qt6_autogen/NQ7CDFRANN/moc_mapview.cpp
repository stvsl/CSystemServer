/****************************************************************************
** Meta object code from reading C++ file 'mapview.h'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.15.2)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include "../../../csystemmain/mapview.h"
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'mapview.h' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.15.2. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_MapView_t {
    QByteArrayData data[11];
    char stringdata0[153];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_MapView_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_MapView_t qt_meta_stringdata_MapView = {
    {
QT_MOC_LITERAL(0, 0, 7), // "MapView"
QT_MOC_LITERAL(1, 8, 9), // "setCenter"
QT_MOC_LITERAL(2, 18, 0), // ""
QT_MOC_LITERAL(3, 19, 10), // "setPoiText"
QT_MOC_LITERAL(4, 30, 10), // "setPoiIcon"
QT_MOC_LITERAL(5, 41, 7), // "searchP"
QT_MOC_LITERAL(6, 49, 3), // "str"
QT_MOC_LITERAL(7, 53, 10), // "setMapType"
QT_MOC_LITERAL(8, 64, 24), // "on_PositionReset_clicked"
QT_MOC_LITERAL(9, 89, 30), // "on_mapSearcher_editingFinished"
QT_MOC_LITERAL(10, 120, 32) // "on_node_Searcher_editingFinished"

    },
    "MapView\0setCenter\0\0setPoiText\0setPoiIcon\0"
    "searchP\0str\0setMapType\0on_PositionReset_clicked\0"
    "on_mapSearcher_editingFinished\0"
    "on_node_Searcher_editingFinished"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_MapView[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       8,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       5,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   54,    2, 0x06 /* Public */,
       3,    1,   55,    2, 0x06 /* Public */,
       4,    1,   58,    2, 0x06 /* Public */,
       5,    1,   61,    2, 0x06 /* Public */,
       7,    0,   64,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       8,    0,   65,    2, 0x08 /* Private */,
       9,    0,   66,    2, 0x08 /* Private */,
      10,    0,   67,    2, 0x08 /* Private */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, QMetaType::Bool,    2,
    QMetaType::Void, QMetaType::Bool,    2,
    QMetaType::Void, QMetaType::QString,    6,
    QMetaType::Void,

 // slots: parameters
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void,

       0        // eod
};

void MapView::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<MapView *>(_o);
        (void)_t;
        switch (_id) {
        case 0: _t->setCenter(); break;
        case 1: _t->setPoiText((*reinterpret_cast< bool(*)>(_a[1]))); break;
        case 2: _t->setPoiIcon((*reinterpret_cast< bool(*)>(_a[1]))); break;
        case 3: _t->searchP((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->setMapType(); break;
        case 5: _t->on_PositionReset_clicked(); break;
        case 6: _t->on_mapSearcher_editingFinished(); break;
        case 7: _t->on_node_Searcher_editingFinished(); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (MapView::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&MapView::setCenter)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (MapView::*)(bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&MapView::setPoiText)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (MapView::*)(bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&MapView::setPoiIcon)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (MapView::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&MapView::searchP)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (MapView::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&MapView::setMapType)) {
                *result = 4;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject MapView::staticMetaObject = { {
    QMetaObject::SuperData::link<QWidget::staticMetaObject>(),
    qt_meta_stringdata_MapView.data,
    qt_meta_data_MapView,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *MapView::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *MapView::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_MapView.stringdata0))
        return static_cast<void*>(this);
    return QWidget::qt_metacast(_clname);
}

int MapView::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QWidget::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 8)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 8;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 8)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 8;
    }
    return _id;
}

// SIGNAL 0
void MapView::setCenter()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void MapView::setPoiText(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void MapView::setPoiIcon(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void MapView::searchP(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void MapView::setMapType()
{
    QMetaObject::activate(this, &staticMetaObject, 4, nullptr);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
