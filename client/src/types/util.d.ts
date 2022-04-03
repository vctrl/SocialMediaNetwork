import { FormItemRule } from 'naive-ui';

export type RulesFor<T> = Record<keyof T, FormItemRule | FormItemRule[]>;
