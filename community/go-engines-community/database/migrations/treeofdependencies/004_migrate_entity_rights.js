(function () {
    db.default_rights.deleteMany({
        "_id": {"$in": [
            "api_entity_read","api_entity_update","api_entity_delete","api_watcher"]}
    });

    db.default_rights.insertMany([
        {
            _id: "api_entity",
            crecord_name: "api_entity",
            crecord_type: "action",
            desc: "Entity",
            type: "CRUD"
        },
        {
            _id: "api_entityservice",
            crecord_name: "api_entityservice",
            crecord_type: "action",
            desc: "Entity service",
            type: "CRUD"
        }
    ]);
    db.default_rights.find({crecord_type: "role"}).forEach(function (doc) {
        var set = {};
        Object.keys(doc.rights).forEach(function (right) {
            switch (right) {
                case "api_watcher":
                    set["rights.api_entityservice"] = doc.rights[right];
                    break;
                case "api_entity_read":
                    if (set["rights.api_entity"]) {
                        set["rights.api_entity"].checksum = set["rights.api_entity"].checksum | 4;
                    } else {
                        set["rights.api_entity"] = {
                            checksum: 4,
                            crecord_type: "right"
                        };
                    }
                    break
                case "api_entity_update":
                    if (set["rights.api_entity"]) {
                        set["rights.api_entity"].checksum = set["rights.api_entity"].checksum | 10;
                    } else {
                        set["rights.api_entity"] = {
                            checksum: 10,
                            crecord_type: "right"
                        };
                    }
                    break
                case "api_entity_delete":
                    if (set["rights.api_entity"]) {
                        set["rights.api_entity"].checksum = set["rights.api_entity"].checksum | 1;
                    } else {
                        set["rights.api_entity"] = {
                            checksum: 1,
                            crecord_type: "right"
                        };
                    }
                    break
            }
        })
        if (Object.keys(set).length > 0) {
            db.default_rights.updateOne(
                {_id: doc._id},
                {
                    $set: set,
                    $unset: {
                        "rights.api_entity_read": "",
                        "rights.api_entity_update": "",
                        "rights.api_entity_delete": "",
                        "rights.api_watcher": ""
                    }
                }
            )
        }
    })
})()
