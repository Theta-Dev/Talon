function formatDate(dateString: string): string {
    return new Date(dateString).toLocaleString(
        /* global navigator */
        navigator.language
    )
}

export {formatDate}
