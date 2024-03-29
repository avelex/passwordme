// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

export function CreatePassfile(arg1:string):Promise<void>;

export function DeletePassfile():Promise<void>;

export function GeneratePassword(arg1:string,arg2:string,arg3:string,arg4:Array<string>):Promise<string>;

export function ImportPassfile():Promise<void>;

export function ListPasswords():Promise<Array<string>>;

export function OpenPassfile(arg1:string):Promise<void>;

export function PassfileExists():Promise<boolean>;

export function PassfileOpened():Promise<boolean>;

export function SavePassword(arg1:string,arg2:string):Promise<void>;

export function ShowPassword(arg1:string):Promise<string>;

export function Version():Promise<string>;
