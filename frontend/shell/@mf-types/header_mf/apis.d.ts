
    export type RemoteKeys = 'header_mf/Header';
    type PackageType<T> = T extends 'header_mf/Header' ? typeof import('header_mf/Header') :any;