export type Move = {
    ID: number;
    Name: string;
    Class: string;
    Type: string;
    PP: number;
    Accuracy: number;
    Priority: number;
    Power: number;
    Description: string;
}

export const newMove = (): Move => ({
    ID: 0,
    Name: "",
    Class: "",
    Type: "",
    PP: 0,
    Accuracy: 0,
    Priority: 0,
    Power: 0,
    Description: "",
})