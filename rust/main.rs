include!("consts.rs");

fn tk_byte(b: &[u8], i: i32) -> i32 {
    if i < 0 || i as usize >= b.len() {
        0
    } else {
        b[i as usize] as i32
    }
}

fn tk_is_space(c: i32) -> bool {
    c == 32 || (9..=13).contains(&c)
}
fn tk_is_digit(c: i32) -> bool {
    (b'0' as i32) <= c && c <= (b'9' as i32)
}
fn tk_is_xdigit(c: i32) -> bool {
    tk_is_digit(c) || ((b'a' as i32) <= c && c <= (b'f' as i32)) || ((b'A' as i32) <= c && c <= (b'F' as i32))
}
fn tk_is_idchar(c: i32) -> bool {
    if c >= 0x80 {
        return true;
    }
    if tk_is_digit(c) {
        return true;
    }
    if (b'A' as i32) <= c && c <= (b'Z' as i32) {
        return true;
    }
    if (b'a' as i32) <= c && c <= (b'z' as i32) {
        return true;
    }
    c == (b'_' as i32) || c == (b'$' as i32)
}

fn tk_class(c: i32) -> i32 {
    if c == 0 {
        return CC_NUL;
    }
    if c == 0xEF {
        return CC_BOM;
    }
    if c >= 0x80 {
        return CC_ID;
    }
    if c == (b'x' as i32) || c == (b'X' as i32) {
        return CC_X;
    }
    if ((b'A' as i32) <= c && c <= (b'W' as i32)) || ((b'a' as i32) <= c && c <= (b'w' as i32)) {
        return CC_KYWD0;
    }
    if ((b'Y' as i32) <= c && c <= (b'Z' as i32)) || ((b'y' as i32) <= c && c <= (b'z' as i32)) || c == (b'_' as i32) {
        return CC_KYWD;
    }
    if tk_is_digit(c) {
        return CC_DIGIT;
    }
    if c == (b'$' as i32) {
        return CC_DOLLAR;
    }
    if c == (b'@' as i32) || c == (b'#' as i32) || c == (b':' as i32) {
        return CC_VARALPHA;
    }
    if c == (b'?' as i32) {
        return CC_VARNUM;
    }
    if tk_is_space(c) {
        if c == 11 {
            return CC_ILLEGAL;
        }
        return CC_SPACE;
    }
    if c == (b'"' as i32) || c == 39 || c == (b'`' as i32) {
        return CC_QUOTE;
    }
    if c == (b'[' as i32) {
        return CC_QUOTE2;
    }
    if c == (b'|' as i32) {
        return CC_PIPE;
    }
    if c == (b'-' as i32) {
        return CC_MINUS;
    }
    if c == (b'<' as i32) {
        return CC_LT;
    }
    if c == (b'>' as i32) {
        return CC_GT;
    }
    if c == (b'=' as i32) {
        return CC_EQ;
    }
    if c == (b'!' as i32) {
        return CC_BANG;
    }
    if c == (b'/' as i32) {
        return CC_SLASH;
    }
    if c == (b'(' as i32) {
        return CC_LP;
    }
    if c == (b')' as i32) {
        return CC_RP;
    }
    if c == (b';' as i32) {
        return CC_SEMI;
    }
    if c == (b'+' as i32) {
        return CC_PLUS;
    }
    if c == (b'*' as i32) {
        return CC_STAR;
    }
    if c == (b'%' as i32) {
        return CC_PERCENT;
    }
    if c == (b',' as i32) {
        return CC_COMMA;
    }
    if c == (b'&' as i32) {
        return CC_AND;
    }
    if c == (b'~' as i32) {
        return CC_TILDA;
    }
    if c == (b'.' as i32) {
        return CC_DOT;
    }
    CC_ILLEGAL
}

fn tk_id_len(b: &[u8], pos: i32) -> i32 {
    let mut j = pos + 1;
    while tk_is_idchar(tk_byte(b, j)) {
        j += 1;
    }
    j - pos
}

fn tk_scan_digits(b: &[u8], mut j: i32, mut t: i32) -> (i32, i32) {
    loop {
        let c = tk_byte(b, j);
        if tk_is_digit(c) {
            j += 1;
        } else if c == (b'_' as i32) {
            t = TK_QNUMBER;
            j += 1;
        } else {
            break;
        }
    }
    (j, t)
}

fn tk_number(b: &[u8], pos: i32) -> (i32, i32) {
    let mut t = TK_INTEGER;
    let mut j;
    if tk_byte(b, pos) == (b'0' as i32)
        && (tk_byte(b, pos + 1) == (b'x' as i32) || tk_byte(b, pos + 1) == (b'X' as i32))
        && tk_is_xdigit(tk_byte(b, pos + 2))
    {
        j = pos + 3;
        loop {
            let c = tk_byte(b, j);
            if tk_is_xdigit(c) {
                j += 1;
            } else if c == (b'_' as i32) {
                t = TK_QNUMBER;
                j += 1;
            } else {
                break;
            }
        }
    } else {
        let r = tk_scan_digits(b, pos, t);
        j = r.0;
        t = r.1;
        if tk_byte(b, j) == (b'.' as i32) {
            if t == TK_INTEGER {
                t = TK_FLOAT;
            }
            j += 1;
            let r = tk_scan_digits(b, j, t);
            j = r.0;
            t = r.1;
        }
        let e = tk_byte(b, j);
        if (e == (b'e' as i32) || e == (b'E' as i32))
            && (tk_is_digit(tk_byte(b, j + 1))
                || ((tk_byte(b, j + 1) == (b'+' as i32) || tk_byte(b, j + 1) == (b'-' as i32))
                    && tk_is_digit(tk_byte(b, j + 2))))
        {
            if t == TK_INTEGER {
                t = TK_FLOAT;
            }
            j += 2;
            let r = tk_scan_digits(b, j, t);
            j = r.0;
            t = r.1;
        }
    }
    while tk_is_idchar(tk_byte(b, j)) {
        t = TK_ILLEGAL;
        j += 1;
    }
    (t, j - pos)
}

fn tk_variable(b: &[u8], pos: i32) -> (i32, i32) {
    let mut t = TK_VARIABLE;
    let mut cnt = 0;
    let mut j = pos + 1;
    let mut done = false;
    while !done {
        let c = tk_byte(b, j);
        if c == 0 {
            done = true;
        } else if tk_is_idchar(c) {
            cnt += 1;
            j += 1;
        } else if c == (b'(' as i32) && cnt > 0 {
            j += 1;
            loop {
                let d = tk_byte(b, j);
                if d == 0 || tk_is_space(d) || d == (b')' as i32) {
                    break;
                }
                j += 1;
            }
            if tk_byte(b, j) == (b')' as i32) {
                j += 1;
            } else {
                t = TK_ILLEGAL;
            }
            done = true;
        } else if c == (b':' as i32) && tk_byte(b, j + 1) == (b':' as i32) {
            j += 2;
        } else {
            done = true;
        }
    }
    if cnt == 0 {
        t = TK_ILLEGAL;
    }
    (t, j - pos)
}

fn tk_scan(b: &[u8], pos: i32, kw: &std::collections::HashMap<&'static [u8], i32>) -> (i32, i32) {
    let c0 = tk_byte(b, pos);
    let cls = tk_class(c0);
    if cls == CC_SPACE {
        let mut j = pos + 1;
        while tk_is_space(tk_byte(b, j)) {
            j += 1;
        }
        return (TK_SPACE, j - pos);
    }
    if cls == CC_BOM {
        if tk_byte(b, pos + 1) == 0xBB && tk_byte(b, pos + 2) == 0xBF {
            return (TK_SPACE, 3);
        }
        return (TK_ID, tk_id_len(b, pos));
    }
    if cls == CC_NUL {
        return (TK_ILLEGAL, 0);
    }
    if cls == CC_MINUS {
        if tk_byte(b, pos + 1) == (b'-' as i32) {
            let mut j = pos + 2;
            while tk_byte(b, j) != (b'\n' as i32) && tk_byte(b, j) != 0 {
                j += 1;
            }
            return (TK_COMMENT, j - pos);
        }
        if tk_byte(b, pos + 1) == (b'>' as i32) {
            let mut n = 2;
            if tk_byte(b, pos + 2) == (b'>' as i32) {
                n = 3;
            }
            return (TK_PTR, n);
        }
        return (TK_MINUS, 1);
    }
    if cls == CC_LP {
        return (TK_LP, 1);
    }
    if cls == CC_RP {
        return (TK_RP, 1);
    }
    if cls == CC_SEMI {
        return (TK_SEMI, 1);
    }
    if cls == CC_PLUS {
        return (TK_PLUS, 1);
    }
    if cls == CC_STAR {
        return (TK_STAR, 1);
    }
    if cls == CC_PERCENT {
        return (TK_REM, 1);
    }
    if cls == CC_COMMA {
        return (TK_COMMA, 1);
    }
    if cls == CC_AND {
        return (TK_BITAND, 1);
    }
    if cls == CC_TILDA {
        return (TK_BITNOT, 1);
    }
    if cls == CC_SLASH {
        if tk_byte(b, pos + 1) != (b'*' as i32) || tk_byte(b, pos + 2) == 0 {
            return (TK_SLASH, 1);
        }
        let mut j = pos + 2;
        while tk_byte(b, j) != 0 {
            if tk_byte(b, j) == (b'*' as i32) && tk_byte(b, j + 1) == (b'/' as i32) {
                return (TK_COMMENT, j + 2 - pos);
            }
            j += 1;
        }
        return (TK_COMMENT, j - pos);
    }
    if cls == CC_EQ {
        if tk_byte(b, pos + 1) == (b'=' as i32) {
            return (TK_EQ, 2);
        }
        return (TK_EQ, 1);
    }
    if cls == CC_LT {
        let c1 = tk_byte(b, pos + 1);
        if c1 == (b'=' as i32) {
            return (TK_LE, 2);
        }
        if c1 == (b'>' as i32) {
            return (TK_NE, 2);
        }
        if c1 == (b'<' as i32) {
            return (TK_LSHIFT, 2);
        }
        return (TK_LT, 1);
    }
    if cls == CC_GT {
        let c1 = tk_byte(b, pos + 1);
        if c1 == (b'=' as i32) {
            return (TK_GE, 2);
        }
        if c1 == (b'>' as i32) {
            return (TK_RSHIFT, 2);
        }
        return (TK_GT, 1);
    }
    if cls == CC_BANG {
        if tk_byte(b, pos + 1) != (b'=' as i32) {
            return (TK_ILLEGAL, 1);
        }
        return (TK_NE, 2);
    }
    if cls == CC_PIPE {
        if tk_byte(b, pos + 1) != (b'|' as i32) {
            return (TK_BITOR, 1);
        }
        return (TK_CONCAT, 2);
    }
    if cls == CC_QUOTE {
        let delim = c0;
        let mut j = pos + 1;
        let mut closed = false;
        while tk_byte(b, j) != 0 {
            if tk_byte(b, j) == delim {
                if tk_byte(b, j + 1) == delim {
                    j += 2;
                } else {
                    closed = true;
                    break;
                }
            } else {
                j += 1;
            }
        }
        if closed {
            if delim == 39 {
                return (TK_STRING, j + 1 - pos);
            }
            return (TK_ID, j + 1 - pos);
        }
        return (TK_ILLEGAL, j - pos);
    }
    if cls == CC_QUOTE2 {
        let mut j = pos + 1;
        while tk_byte(b, j) != 0 && tk_byte(b, j) != (b']' as i32) {
            j += 1;
        }
        if tk_byte(b, j) == (b']' as i32) {
            return (TK_ID, j + 1 - pos);
        }
        return (TK_ILLEGAL, j - pos);
    }
    if cls == CC_DOT {
        if !tk_is_digit(tk_byte(b, pos + 1)) {
            return (TK_DOT, 1);
        }
        return tk_number(b, pos);
    }
    if cls == CC_DIGIT {
        return tk_number(b, pos);
    }
    if cls == CC_VARNUM {
        let mut j = pos + 1;
        while tk_is_digit(tk_byte(b, j)) {
            j += 1;
        }
        return (TK_VARIABLE, j - pos);
    }
    if cls == CC_DOLLAR || cls == CC_VARALPHA {
        return tk_variable(b, pos);
    }
    if cls == CC_KYWD0 {
        if tk_class(tk_byte(b, pos + 1)) > CC_KYWD {
            return (TK_ID, tk_id_len(b, pos));
        }
        let mut j = pos + 2;
        while tk_class(tk_byte(b, j)) <= CC_KYWD && tk_byte(b, j) != 0 {
            j += 1;
        }
        if tk_is_idchar(tk_byte(b, j)) {
            return (TK_ID, tk_id_len(b, pos));
        }
        let mut lw: Vec<u8> = b[pos as usize..j as usize].to_vec();
        for x in lw.iter_mut() {
            if (b'A'..=b'Z').contains(x) {
                *x += 32;
            }
        }
        if let Some(v) = kw.get(lw.as_slice()) {
            return (*v, j - pos);
        }
        return (TK_ID, j - pos);
    }
    if cls == CC_X {
        if tk_byte(b, pos + 1) == 39 {
            let mut j = pos + 2;
            while tk_is_xdigit(tk_byte(b, j)) {
                j += 1;
            }
            let mut bad = false;
            if tk_byte(b, j) != 39 || (j - pos) % 2 == 1 {
                bad = true;
            }
            if bad {
                while tk_byte(b, j) != 0 && tk_byte(b, j) != 39 {
                    j += 1;
                }
            }
            if tk_byte(b, j) != 0 {
                j += 1;
            }
            if bad {
                return (TK_ILLEGAL, j - pos);
            }
            return (TK_BLOB, j - pos);
        }
        return (TK_ID, tk_id_len(b, pos));
    }
    if cls == CC_KYWD || cls == CC_ID {
        return (TK_ID, tk_id_len(b, pos));
    }
    (TK_ILLEGAL, 1)
}

fn main() {
    let data = std::fs::read("workload.sql").expect("cannot open workload.sql");
    let mut lines: Vec<&[u8]> = data.split(|&c| c == b'\n').collect();
    if let Some(last) = lines.last() {
        if last.is_empty() {
            lines.pop();
        }
    }
    let nlines = lines.len();
    let kw = kw_map();
    let mut count: i64 = 0;
    let mut tsum: i64 = 0;
    let mut lsum: i64 = 0;
    for _ in 0..20 {
        for line in &lines {
            let mut pos: i32 = 0;
            while (pos as usize) < line.len() {
                let (t, n) = tk_scan(line, pos, &kw);
                if n == 0 {
                    break;
                }
                count += 1;
                tsum += t as i64;
                lsum += n as i64;
                pos += n;
            }
        }
    }
    println!(
        "Rust: {} tokens, typesum={} lensum={} (20 passes, {} lines)",
        count, tsum, lsum, nlines
    );
}
