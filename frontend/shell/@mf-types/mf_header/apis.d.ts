
    export type RemoteKeys = 'mf_header/Header';
    type PackageType<T> = T extends 'mf_header/Header' ? typeof import('mf_header/Header') :any;