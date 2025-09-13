import React, { type ButtonHTMLAttributes, type MouseEvent, type ReactNode } from "react";

export type ButtonType = "warn" | "success" | "neutral" | "primary" | "error" | "disabled";
export type ButtonVariant = "solid" | "outline";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    children: ReactNode;
    buttonType?: ButtonType;
    buttonVariant?: ButtonVariant;
    onClick?: (e: MouseEvent<HTMLButtonElement>) => Promise<void> | void;
    isLoading?: boolean;
};

const baseClasses = "rounded-lg box-border font-main text-small duration-200 hover:shadow-xl active:shadow-none";
const solidBaseClasses = "px-[20px] py-[9px]"
const outlineBaseClasses = "border-2 hover:border-transparent px-[18px] py-[7px]"

const buttonVariantsClasses = {
    warn: {
        solid: "bg-warn text-primary hover:bg-warn-hl",
        outline: "border-warn text-warn hover:bg-warn-hl hover:text-primary",
    },
    success: {
        solid: "bg-success text-white hover:bg-success-hl",
        outline: "border-success text-success hover:bg-success-hl hover:text-white",
    },
    neutral: {
        solid: "bg-neutral text-white hover:bg-neutral-hl",
        outline: "border-neutral text-neutral hover:bg-neutral-hl hover:text-white",
    },
    primary: {
        solid: "bg-primary text-white hover:bg-primary-hl",
        outline: "border-primary text-primary hover:bg-primary hover:text-white",
    },
    error: {
        solid: "bg-error text-white hover:bg-error-hl",
        outline: "border-error text-error hover:bg-error-hl hover:text-white",
    },
    disabled: {
        solid: "bg-secondary text-tertiary cursor-not-allowed hover:shadow-none!",
        outline: "bg-secondary text-tertiary cursor-not-allowed hover:shadow-none!",
    }
}

const Button: React.FC<ButtonProps> = ({
    children,
    buttonType = "neutral",
    buttonVariant = "solid",
    onClick,
    isLoading = false,
    disabled = false,
    className = "",
    ...props
}: ButtonProps): React.JSX.Element => {

    const handleClick = async (e: MouseEvent<HTMLButtonElement>) => {
        if (onClick) {
            try {
                await onClick(e);
            } catch (error) {
                console.error(error);
            }
        }
    };

    const variantBaseClasses = buttonVariant === "solid" ? solidBaseClasses : outlineBaseClasses;
    const variantClasses = disabled || isLoading ? buttonVariantsClasses["disabled"][buttonVariant] : buttonVariantsClasses[buttonType][buttonVariant];

    return <button className={`${baseClasses} ${variantBaseClasses} ${variantClasses} ${className}`} onClick={handleClick} disabled={isLoading || disabled} {...props}>
        {isLoading ? "Loading..." : children}
    </button>
}

export default Button;

