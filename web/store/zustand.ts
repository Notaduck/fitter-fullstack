import { create } from 'zustand';

type ActivityStore = {
    recordId: number | null;
    recordValue: number | null;
    recordPosition: number | null; // Adjusted to match the name in the create function
    setRecordId: (newId: number) => void;
    setRecordValue: (newValue: number) => void;
    setRecordPosition: (newPos: number) => void;
};

export const useActivityStore = create<ActivityStore>((set) => ({
    recordId: null,
    recordPosition: null, // Adjusted to match the name in the type definition
    recordValue: null,
    setRecordId: (newId: number) => set((state) => ({ recordId: newId })),
    setRecordValue: (newValue: number) => set((state) => ({ recordValue: newValue })),
    setRecordPosition: (newPos: number) => set((state) => ({ recordPosition: newPos })),
}));
