INSERT INTO pages (name) VALUES 
    ('home')
;

INSERT INTO widgets (mf_name, component_name, entry_url) VALUES
    ('mf-header',    'Header',   'http://localhost:5001/remoteEntry.js'),
    ('mf-footer',    'Footer',   'http://localhost:5002/remoteEntry.js'),
    ('mf-items-list', 'ItemsList', 'http://localhost:5003/remoteEntry.js')
;

INSERT INTO pages_widgets (level, page_id, widget_id) VALUES 
    (1, 1, 1),
    (2, 1, 3),
    (3, 1, 2)
;